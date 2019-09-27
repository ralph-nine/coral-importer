//go:generate easyjson -all models.go
package livefyre

import (
	"fmt"

	"gitlab.com/coralproject/coral-importer/common/coral"
)

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
func TranslateComment(tenantID string, in *Comment) *coral.Comment {
	comment := coral.NewComment(tenantID)
	comment.ID = fmt.Sprintf("%d", in.ID)
	if in.ParentID > 0 {
		comment.ParentID = fmt.Sprintf("%d", in.ParentID)
		comment.ParentRevisionID = comment.ParentID
	}
	comment.AuthorID = in.AuthorID
	comment.ActionCounts = map[string]int{}
	comment.Tags = []coral.CommentTag{}
	comment.CreatedAt.Time = in.Created.Time

	switch in.State {
	case 0:
		comment.Status = "REJECTED"
	case 1:
		comment.Status = "APPROVED"
	case 2:
		comment.Status = "REJECTED"
	case 3:
		comment.Status = "NONE"
	case 4:
		comment.Status = "PREMOD"
	case 5:
		comment.Status = "REJECTED"
	default:
		comment.Status = "NONE"
	}

	revision := coral.Revision{
		ID:           comment.ID,
		Body:         coral.HTML(in.BodyHTML),
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
	Title    string    `json:"title"`
	Source   string    `json:"source" validate:"required,url"`
	Comments []Comment `json:"comments" validate:"required"`
	Created  Time      `json:"created"  validate:"required"`
}

// TranslateStory will copy over simple fields to the new coral.Story.
func TranslateStory(tenantID string, in *Story) *coral.Story {
	story := coral.NewStory(tenantID)
	story.ID = in.ID
	story.URL = in.Source
	story.Metadata.Title = in.Title
	story.CreatedAt.Time = in.Created.Time

	return story
}
