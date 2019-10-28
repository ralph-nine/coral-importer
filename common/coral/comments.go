//go:generate easyjson -all comments.go
package coral

// RevisionMetadata is the metadata associated with a given Revision for a
// Comment in Coral.
type RevisionMetadata struct{}

// Revision is a given revision of a Comment in Coral.
type Revision struct {
	ID           string           `json:"id" conform:"trim" validate:"required"`
	Body         HTML             `json:"body" conform:"trim" validate:"required"`
	ActionCounts map[string]int   `json:"actionCounts" validate:"required"`
	Metadata     RevisionMetadata `json:"metadata" validate:"required"`
	CreatedAt    Time             `json:"createdAt" validate:"required"`
}

// CommentTag is a Tag associated with a Comment in Coral.
type CommentTag struct {
	Type      string `json:"type" conform:"trim" validate:"oneof=STAFF FEATURED,required"`
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedAt Time   `json:"createdAt" validate:"required"`
}

// Comment is the base Coral Comment that is used in Coral.
type Comment struct {
	TenantID         string         `json:"tenantID" validate:"required"`
	ID               string         `json:"id" conform:"trim" validate:"required"`
	AncestorIDs      []string       `json:"ancestorIDs" validate:"required"`
	ParentID         string         `json:"parentID,omitempty" conform:"trim"`
	ParentRevisionID string         `json:"parentRevisionID,omitempty" conform:"trim"`
	AuthorID         string         `json:"authorID" conform:"trim" validate:"required"`
	StoryID          string         `json:"storyID" conform:"trim" validate:"required"`
	Revisions        []Revision     `json:"revisions" validate:"required"`
	Status           string         `json:"status" conform:"trim" validate:"oneof=NONE APPROVED REJECTED PREMOD SYSTEM_WITHHELD,required"`
	ActionCounts     map[string]int `json:"actionCounts" validate:"required"`
	ChildIDs         []string       `json:"childIDs" validate:"required"`
	Tags             []CommentTag   `json:"tags" validate:"required"`
	ChildCount       int            `json:"childCount" validate:"gte=0"`
	CreatedAt        Time           `json:"createdAt" validate:"required"`
	DeletedAt        *Time          `json:"deletedAt,omitempty"`
	Imported         bool           `json:"imported"`
}

// NewComment will return an initialized Comment.
func NewComment(tenantID string) *Comment {
	return &Comment{
		TenantID:     tenantID,
		AncestorIDs:  []string{},
		Revisions:    []Revision{},
		ActionCounts: map[string]int{},
		ChildIDs:     []string{},
		Tags:         []CommentTag{},
		Imported:     true,
	}
}
