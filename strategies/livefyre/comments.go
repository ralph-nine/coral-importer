package livefyre

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gitlab.com/coralproject/coral-importer/common"
	"gitlab.com/coralproject/coral-importer/common/coral"
	"gitlab.com/coralproject/coral-importer/common/importer"
)

// Time is the time.Time representation that LiveFyre uses.
type Time struct {
	time.Time
}

// UnmarshalJSON is the custom unmarshaler for the input JSON from LiveFyre.
func (t *Time) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse("2006-01-02T15:04:05", strings.Trim(string(buf), `"`))
	if err != nil {
		return err
	}

	t.Time = tt

	return nil
}

// Comment is the Comment as exported from the LiveFyre platform.
type Comment struct {
	ID       int    `json:"id" validate:"required"`
	BodyHTML string `json:"body_html" validate:"required"`
	ParentID int    `json:"parent_id"`
	AuthorID string `json:"author_id"`
	State    int    `json:"state"`
	Created  Time   `json:"created" validate:"required"`
}

// TranslateComment will copy over simple fields to the new coral.Comment.
func TranslateComment(in *Comment) *coral.Comment {
	comment := coral.NewComment()
	comment.ID = fmt.Sprintf("%d", in.ID)
	if in.ParentID > 0 {
		comment.ParentID = fmt.Sprintf("%d", in.ParentID)
		comment.ParentRevisionID = comment.ParentID
	}
	comment.AuthorID = in.AuthorID
	comment.CreatedAt.Time = in.Created.Time

	switch in.State {
	// TODO: implement
	default:
		comment.Status = "NONE"
	}

	revision := coral.Revision{
		ID:           comment.ID,
		Body:         in.BodyHTML,
		Metadata:     coral.RevisionMetadata{},
		ActionCounts: map[string]int{},
	}
	revision.CreatedAt.Time = in.Created.Time

	comment.Revisions = append(comment.Revisions, revision)

	return comment
}

// Story is the Story as exported from the LiveFyre platform.
type Story struct {
	ID       string    `json:"id" validate:"required"`
	Source   string    `json:"source" validate:"required,url"`
	Comments []Comment `json:"comments" validate:"required"`
	Created  Time      `json:"created"  validate:"required"`
}

// TranslateStory will copy over simple fields to the new coral.Story.
func TranslateStory(in *Story) *coral.Story {
	story := coral.NewStory()
	story.ID = in.ID
	story.URL = in.Source
	story.CreatedAt = in.Created.Time

	return story
}

// Comments will handle a data import task for importing comments into Coral
// from a LiveFyre export.
func Comments(c *cli.Context) error {
	// Grab the debug mode.
	debug := c.GlobalBool("debug")
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// Grab the input fileName.
	input := c.String("input")

	// Open that file for reading.
	f, err := os.Open(input)
	if err != nil {
		return errors.Wrap(err, "could not open --input for reading")
	}
	defer f.Close()

	logrus.WithField("input", input).Info("opened for reading")

	// Setup the scanner.
	r := bufio.NewReader(f)

	// Setup the importers.
	stories := importer.New("stories")
	comments := importer.New("comments")

	// Configure the context for the importers.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the importers.
	var wg sync.WaitGroup
	wg.Add(2)
	go comments.Start(ctx, &wg)
	go stories.Start(ctx, &wg)

	// Keep track of the processed lines.
	lines := 0

	// Start reading the stories line by line from the file.
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrap(err, "couldn't read the file")
		}

		// Parse the Story from the file.
		var in Story
		if err := json.Unmarshal([]byte(line), &in); err != nil {
			return errors.Wrap(err, "could not parse a comment in the --input file")
		}

		// Increment the line count.
		lines++

		// Check the input to ensure we're validated.
		if err := common.Check(&in); err != nil {
			return errors.Wrap(err, "checking failed input Story")
		}

		// Translate the Story to a coral.Story.
		story := TranslateStory(&in)

		// Check the story to ensure we're validated.
		if err := common.Check(story); err != nil {
			return errors.Wrap(err, "checking failed output coral.Story")
		}

		// Send the story to the importer.
		if err := stories.Import(*story); err != nil {
			return errors.Wrap(err, "failed to import the story")
		}

		logrus.WithFields(logrus.Fields{
			"storyID": story.ID,
			"line":    lines,
		}).Info("imported story")

		// Collect all the stories comments so we can process family
		// relationships as well.
		storyComments := make([]*coral.Comment, 0, len(in.Comments))

		// Reconstruct family relationships for these comments.
		r := common.NewReconstructor()

		// Translate the comments.
		for _, inc := range in.Comments {
			if inc.AuthorID == "" {
				logrus.WithFields(logrus.Fields{
					"storyID":   story.ID,
					"commentID": inc.ID,
					"line":      lines,
				}).Warn("comment was missing author_id field")
				continue
			}

			// Check the comment to ensure we're validated.
			if err := common.Check(&inc); err != nil {
				return errors.Wrapf(err, "checking failed input Comment for Story %s", story.ID)
			}

			// Translate the Comment to a coral.Comment.
			comment := TranslateComment(&inc)
			comment.StoryID = story.ID

			// Check the comment to ensure we're validated.
			if err := common.Check(comment); err != nil {
				return errors.Wrap(err, "checking failed output coral.Comment")
			}

			// Add the comment to the reconstructor.
			r.AddComment(comment)

			// Add it to the story comments.
			storyComments = append(storyComments, comment)
		}

		// Send the comments off to the importer.
		for _, comment := range storyComments {
			// Add reconstructed data.
			comment.ChildIDs = r.GetChildren(comment.ID)
			comment.ChildCount = len(comment.ChildIDs)
			comment.AncestorIDs = r.GetAncestors(comment.ID)

			// Send the comment to the importer.
			if err := comments.Import(*comment); err != nil {
				return errors.Wrap(err, "failed to import the comment")
			}

			logrus.WithFields(logrus.Fields{
				"storyID":   story.ID,
				"commentID": comment.ID,
				"line":      lines,
			}).Info("imported comment")
		}
	}

	// Close the importers and wait till they're done.
	comments.Done()
	stories.Done()
	wg.Wait()

	logrus.WithField("lines", lines).Info("finished processing")

	return nil
}
