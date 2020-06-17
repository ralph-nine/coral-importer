package csv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/common"
	"gitlab.com/coralproject/coral-importer/common/coral"
	"gitlab.com/coralproject/coral-importer/common/pipeline"
)

// Import will perform the actual import process for the CSV strategy.
func Import(c *cli.Context) error {
	// tenantID is the ID of the Tenant that we are importing these documents
	// for.
	tenantID := c.String("tenantID")

	// siteID is the ID of the Site that we're importing records for.
	siteID := c.String("siteID")

	// output is the name of the folder where we are placing our outputted dumps
	// ready for MongoDB import.
	output := c.String("output")

	// input is the name of the folder where we are loading out collections
	// from the MongoDB export.
	input := c.String("input")

	// auth is the identifier for the type of authentication profiles to be
	// created for the users.
	authMode := c.String("auth")
	if authMode != "local" && authMode != "sso" {
		return errors.Errorf("invalid --auth provided \"%s\", only \"sso\" or \"local\" is supported")
	}

	// Validate that the collection files we expect exist in the input folder.
	if err := validateCollectionFilesExist(input); err != nil {
		return err
	}

	// Mark when we started.
	started := time.Now()

	logrus.Debug("starting comment map processing")

	// Write out all the comments to ${output}/comments.csv.
	commentsFileName := filepath.Join(input, "comments.csv")
	commentMap, err := pipeline.NewMapAggregator(
		pipeline.MergeTaskAggregatorOutputPipelines(
			pipeline.FanAggregatingProcessor(
				pipeline.NewCSVFileReader(commentsFileName, CommentColumns),
				ProcessCommentMap(),
			),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("could not process comments")
		return err
	}

	logrus.WithField("comments", len(commentMap["storyID"])).WithField("children", len(commentMap["parentID"])).Debug("finished loading comments into map")

	startedReconstructionAt := time.Now()
	logrus.Debug("reconstructing families based on parent id map")

	// Reconstruct all the family relationships from the parentID map.
	reconstructor := common.NewReconstructor()
	for commentID, parentIDs := range commentMap["parentID"] {
		if len(parentIDs) == 1 {
			reconstructor.AddIDs(commentID, parentIDs[0])
		} else {
			reconstructor.AddIDs(commentID, "")
		}
	}

	logrus.WithField("took", time.Since(startedReconstructionAt).String()).Debug("finished family reconstruction")

	// Delete the reference to the parentID map that we don't need any more.
	delete(commentMap, "parentID")

	startedSummerAt := time.Now()
	logrus.Debug("counting comment status")

	// Load all the comment statuses by reading the comments.json file again.
	statusCounts, err := pipeline.NewSummer(
		pipeline.MergeTaskSummerOutputPipelines(
			pipeline.FanSummerProcessor(
				pipeline.NewCSVFileReader(commentsFileName, CommentColumns),
				ProcessCommentStatusMap(),
			),
		),
	)
	if err != nil {
		logrus.WithError(err).Error("could not process status counts")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"took": time.Since(startedSummerAt).String(),
	}).Debug("finished counting comment status")

	startedCommentsAt := time.Now()
	logrus.Debug("processing comments")

	if err := pipeline.NewFileWriter(
		output,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewCSVFileReader(commentsFileName, CommentColumns),
				ProcessComments(tenantID, siteID, reconstructor),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process comments")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"took": time.Since(startedCommentsAt).String(),
	}).Debug("finished processing comments")

	startedUsersAt := time.Now()
	logrus.Debug("processing users")

	// Write out all the users to ${output}/users.csv.
	usersFileName := filepath.Join(input, "users.csv")
	if err := pipeline.NewFileWriter(
		output,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewCSVFileReader(usersFileName, UserColumns),
				ProcessUsers(tenantID, statusCounts, authMode),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process users")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"took": time.Since(startedUsersAt).String(),
	}).Debug("finished processing users")

	startedStoriesAt := time.Now()
	logrus.Debug("processing stories")

	// Write out all the stories to ${output}/stories.csv.
	storiesFileName := filepath.Join(input, "stories.csv")
	if err := pipeline.NewFileWriter(
		output,
		pipeline.MergeTaskWriterOutputPipelines(
			pipeline.FanWritingProcessors(
				pipeline.NewCSVFileReader(storiesFileName, StoryColumns),
				ProcessStories(tenantID, siteID, statusCounts),
			),
		),
	); err != nil {
		logrus.WithError(err).Error("could not process stories")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"took": time.Since(startedStoriesAt).String(),
	}).Debug("finished processing stories")

	// Mark when we finished.
	logrus.WithField("took", time.Since(started).String()).Info("finished processing")

	return nil
}

// IsHeaderRow will return true when the row contains the first field value as
// "id".
func IsHeaderRow(input *pipeline.TaskReaderInput) bool {
	return strings.ToLower(input.Fields[0]) == "id"
}

// ProcessCommentMap will collect maps based on the comment data.
func ProcessCommentMap() pipeline.AggregatingProcessor {
	return func(writer pipeline.AggregationWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		c, err := ParseComment(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to parse comment")
			return nil
		}

		writer("status", c.ID, TranslateCommentStatus(c.Status))
		writer("storyID", c.ID, c.StoryID)
		writer("parentID", c.ID, c.ParentID)

		return nil
	}
}

// ProcessCommentStatusMap will link up comment statuses with the story ID.
func ProcessCommentStatusMap() pipeline.SummerProcessor {
	return func(writer pipeline.SummerWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		c, err := ParseComment(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to parse comment")
			return nil
		}

		status := TranslateCommentStatus(c.Status)

		// Add the status to the map referencing the story id.
		writer("story:"+c.StoryID, status, 1)

		// Add the status to the map referencing the user id.
		writer("user:"+c.AuthorID, status, 1)

		return nil
	}
}

// ProcessComments will emit a comment for every valid CSV line in the input file.
func ProcessComments(tenantID, siteID string, r *common.Reconstructor) pipeline.WritingProcessor {
	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	var p = bluemonday.UGCPolicy()

	return func(write pipeline.CollectionWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		c, err := ParseComment(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to parse comment")
			return nil
		}

		createdAt, err := time.Parse(time.RFC3339, c.CreatedAt)
		if err != nil {
			return errors.Wrap(err, "could not parse created_at")
		}

		comment := coral.NewComment(tenantID, siteID)
		comment.ID = c.ID
		comment.AuthorID = c.AuthorID
		comment.StoryID = c.StoryID
		comment.CreatedAt.Time = createdAt

		rawBody := strings.Replace(c.Body, "\n", "<br/>", -1)
		body := coral.HTML(p.Sanitize(rawBody))

		revision := coral.Revision{
			ID:           comment.ID,
			Body:         body,
			Metadata:     coral.RevisionMetadata{},
			ActionCounts: map[string]int{},
			CreatedAt: coral.Time{
				Time: createdAt,
			},
		}
		comment.Revisions = []coral.Revision{
			revision,
		}
		comment.ParentID = c.ParentID
		comment.Status = TranslateCommentStatus(c.Status)

		// ID of the parent is the same as the revision ID.
		comment.ParentRevisionID = comment.ParentID

		// Add reconstructed data.
		comment.ChildIDs = r.GetChildren(comment.ID)
		comment.ChildCount = len(comment.ChildIDs)
		comment.AncestorIDs = r.GetAncestors(comment.ID)

		if err := common.Check(comment); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to process compiled comment")
			return nil
		}

		if err := write("comments", comment); err != nil {
			return errors.Wrap(err, "couldn't write out comment")
		}

		return nil
	}
}

// ProcessStories will emit a story for every valid CSV line in the input file.
func ProcessStories(tenantID, siteID string, statusCounts map[string]map[string]int) pipeline.WritingProcessor {
	return func(write pipeline.CollectionWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if input.Line == 1 && IsHeaderRow(input) {
			return nil
		}

		s, err := ParseStory(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to parse story")
			return nil
		}

		story := coral.NewStory(tenantID, siteID)
		story.ID = s.ID
		story.URL = s.URL

		// Get the status counts for this story.
		storyStatusCounts, ok := statusCounts["story:"+story.ID]
		if ok {
			story.CommentCounts.Status.Approved = storyStatusCounts["APPROVED"]
			story.CommentCounts.Status.None = storyStatusCounts["NONE"]
			story.CommentCounts.Status.Rejected = storyStatusCounts["REJECTED"]

			// Begin updating the story's moderation queue counts.
			story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.None
			story.CommentCounts.ModerationQueue.Total += story.CommentCounts.Status.Premod
			story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.None
			story.CommentCounts.ModerationQueue.Queues.Unmoderated += story.CommentCounts.Status.Premod
		}

		// Copy over the metadata.
		if s.Title != "" {
			story.Metadata.Title = s.Title
		}
		if s.Author != "" {
			story.Metadata.Author = s.Author
		}
		if s.PublishedAt != "" {
			publishedAt, err := time.Parse(time.RFC3339, s.PublishedAt)
			if err != nil {
				return errors.Wrap(err, "could not parse published_at")
			}

			story.Metadata.PublishedAt = &coral.Time{
				Time: publishedAt,
			}
		}

		story.CreatedAt.Time = story.ImportedAt.Time

		// Copy over the closed at date if provided.
		if s.ClosedAt != "" {
			closedAt, err := time.Parse(time.RFC3339, s.ClosedAt)
			if err != nil {
				return errors.Wrap(err, "could not parse closed_at")
			}

			story.ClosedAt = &coral.Time{
				Time: closedAt,
			}
		}

		if err := common.Check(story); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to process compiled story")
			return nil
		}

		if err := write("stories", story); err != nil {
			return errors.Wrap(err, "couldn't write out story")
		}

		return nil
	}
}

// ProcessUsers will emit a user for every valid CSV line in the input file.
func ProcessUsers(tenantID string, statusCounts map[string]map[string]int, auth string) pipeline.WritingProcessor {
	return func(write pipeline.CollectionWriter, input *pipeline.TaskReaderInput) error {
		// Ensure we skip the line if it's a header line.
		if IsHeaderRow(input) {
			return nil
		}

		u, err := ParseUser(input.Fields)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to parse user")
			return nil
		}

		// Parse the user from the file.
		user := coral.NewUser(tenantID)
		user.ID = u.ID
		user.Email = u.Email

		// Get the status counts for this user.
		userStatusCounts, ok := statusCounts["user:"+user.ID]
		if ok {
			user.CommentCounts.Status.Approved = userStatusCounts["APPROVED"]
			user.CommentCounts.Status.None = userStatusCounts["NONE"]
			user.CommentCounts.Status.Rejected = userStatusCounts["REJECTED"]
		}

		// created_at
		if u.CreatedAt != "" {
			createdAt, err := time.Parse(time.RFC3339, u.CreatedAt)
			if err != nil {
				return errors.Wrap(err, "could not parse created_at")
			}

			user.CreatedAt.Time = createdAt
		} else {
			user.CreatedAt.Time = time.Now()
		}

		// username
		user.Username = u.Username
		user.Status.UsernameStatus.History = []coral.UserUsernameStatusHistory{
			{
				ID:        uuid.NewV1().String(),
				Username:  user.Username,
				CreatedBy: user.ID,
				CreatedAt: user.CreatedAt,
			},
		}

		// role
		user.Role = TranslateUserRole(u.Role)

		// banned
		switch strings.ToLower(u.Banned) {
		case "true":
			user.Status.BanStatus.Active = true
			user.Status.BanStatus.History = []coral.UserBanStatusHistory{
				{
					ID:        uuid.NewV1().String(),
					Message:   "",
					Active:    true,
					CreatedAt: user.CreatedAt,
				},
			}
		case "false":
			fallthrough
		default:
			user.Status.BanStatus.Active = false
		}

		if auth == "local" {
			user.Profiles = []coral.UserProfile{
				{
					ID:         user.Email,
					Type:       "local",
					Password:   uuid.NewV4().String(),
					PasswordID: uuid.NewV4().String(),
				},
			}
		} else if auth == "sso" {
			user.Profiles = []coral.UserProfile{
				{
					ID:           user.ID,
					Type:         "sso",
					LastIssuedAt: &user.CreatedAt,
				},
			}
		}

		// Check the user.
		if err := common.Check(user); err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"line":   input.Line,
				"fields": input.Fields,
			}).Warn("failed to process compiled user")
			return nil
		}

		if err := write("users", user); err != nil {
			return errors.Wrap(err, "couldn't write out user")
		}

		return nil
	}
}

// validateCollectionFilesExist will ensure that all the collection files that
// we expect to be in the input directory actually exist.
func validateCollectionFilesExist(input string) error {
	var collections = []string{
		"users",
		"stories",
		"comments",
	}

	for _, collection := range collections {
		filePath := filepath.Join(input, fmt.Sprintf("%s.csv", collection))
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return errors.Errorf("%s does not exist", filePath)
		}
	}

	return nil
}
