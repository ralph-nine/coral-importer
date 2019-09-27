package livefyre

import (
	easyjson "github.com/mailru/easyjson"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/coralproject/coral-importer/common"
	"gitlab.com/coralproject/coral-importer/common/coral"
	"gitlab.com/coralproject/coral-importer/common/pipeline"
)

func ProcessComments(tenantID string, authorIDs map[string]string) pipeline.WritingProcessor {
	return func(write pipeline.CollectionWriter, n *pipeline.TaskReaderInput) error {

		// Parse the Story from the file.
		var in Story
		if err := easyjson.Unmarshal([]byte(n.Input), &in); err != nil {
			return errors.Wrap(err, "could not parse a comment in the --input file")
		}

		// Check the input to ensure we're validated.
		if err := common.Check(&in); err != nil {
			return errors.Wrap(err, "checking failed input Story")
		}

		// Translate the Story to a coral.Story.
		story := TranslateStory(tenantID, &in)

		// Check the story to ensure we're validated.
		if err := common.Check(story); err != nil {
			return errors.Wrap(err, "checking failed output coral.Story")
		}

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
					"line":      n.Line,
				}).Warn("comment was missing author_id field")
				continue
			}

			// Check the comment to ensure we're validated.
			if err := common.Check(&inc); err != nil {
				return errors.Wrapf(err, "checking failed input Comment for Story %s", story.ID)
			}

			// Remap the authorID.
			authorID := authorIDs[inc.AuthorID]
			if authorID == "" {
				logrus.WithFields(logrus.Fields{
					"storyID":   story.ID,
					"commentID": inc.ID,
					"authorID":  inc.AuthorID,
					"line":      n.Line,
				}).Warn("comment author_id did not exist in author map")
				continue
			}
			inc.AuthorID = authorID

			// Translate the Comment to a coral.Comment.
			comment := TranslateComment(tenantID, &inc)
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
			if err := write("comments", comment); err != nil {
				return errors.Wrap(err, "couldn't write out comment")
			}

			logrus.WithFields(logrus.Fields{
				"storyID":   story.ID,
				"commentID": comment.ID,
				"line":      n.Line,
			}).Debug("imported comment")
		}

		// Increment the stories comment counts.
		for _, comment := range storyComments {
			story.IncrementCommentCounts(comment.Status)
		}

		// Send the story to the importer.
		if err := write("stories", story); err != nil {
			return errors.Wrap(err, "couldn't write out story")
		}

		logrus.WithFields(logrus.Fields{
			"storyID": story.ID,
			"line":    n.Line,
		}).Debug("imported story")

		logrus.WithFields(logrus.Fields{
			"storyID":  story.ID,
			"line":     n.Line,
			"comments": len(storyComments),
		}).Info("finished line")

		return nil
	}
}
