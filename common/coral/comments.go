package coral

import (
	"encoding/json"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"$date": t.Time,
	})
}

// RevisionMetadata is the metadata associated with a given Revision for a
// Comment in Coral.
type RevisionMetadata struct{}

// Revision is a given revision of a Comment in Coral.
type Revision struct {
	ID           string           `json:"id" bson:"id" conform:"trim" validate:"required"`
	Body         string           `json:"body" bson:"body" conform:"trim" validate:"required"`
	ActionCounts map[string]int   `json:"actionCounts" bson:"actionCounts" validate:"required"`
	Metadata     RevisionMetadata `json:"metadata" bson:"metadata" validate:"required"`
	CreatedAt    Time             `json:"createdAt" bson:"createdAt" validate:"required"`
}

// CommentTag is a Tag associated with a Comment in Coral.
type CommentTag struct {
	Type      string `json:"type" bson:"type" conform:"trim" validate:"required"`
	CreatedBy string `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	CreatedAt Time   `json:"createdAt" bson:"createdAt" validate:"required"`
}

// Comment is the base Coral Comment that is used in Coral.
type Comment struct {
	ID               string         `json:"id" bson:"id" conform:"trim" validate:"required"`
	AncestorIDs      []string       `json:"ancestorIDs" bson:"ancestorIDs" validate:"required"`
	ParentID         string         `json:"parentID,omitempty" bson:"parentID" conform:"trim"`
	ParentRevisionID string         `json:"parentRevisionID,omitempty" bson:"parentRevisionID,omitempty" conform:"trim"`
	AuthorID         string         `json:"authorID" bson:"authorID" conform:"trim" validate:"required"`
	StoryID          string         `json:"storyID" bson:"storyID" conform:"trim" validate:"required"`
	Revisions        []Revision     `json:"revisions" bson:"revisions" validate:"required"`
	Status           string         `json:"status" bson:"status" conform:"trim" validate:"oneof=NONE APPROVED REJECTED PREMOD SYSTEM_WITHHELD,required"`
	ActionCounts     map[string]int `json:"actionCounts" bson:"actionCounts" validate:"required"`
	ChildIDs         []string       `json:"childIDs" bson:"childIDs" validate:"required"`
	Tags             []CommentTag   `json:"tags" bson:"tags" validate:"required"`
	ChildCount       int            `json:"childCount" bson:"childCount" validate:"gte=0"`
	CreatedAt        Time           `json:"createdAt" bson:"createdAt" validate:"required"`
	Imported         bool           `json:"imported" bson:"imported"`
}

// NewComment will return an initialized Comment.
func NewComment() *Comment {
	return &Comment{
		AncestorIDs:  []string{},
		Revisions:    []Revision{},
		ActionCounts: map[string]int{},
		ChildIDs:     []string{},
		Tags:         []CommentTag{},
		Imported:     true,
	}
}
