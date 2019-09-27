package coral

import "time"

// Story is thye base Coral Story that is used in Coral.
type Story struct {
	ID        string    `json:"id" bson:"id" conform:"trim" validate:"required"`
	URL       string    `json:"url" bson:"url" validate:"required,url"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt" validate:"required"`
}

// NewStory will return an initalized Story.
func NewStory() *Story {
	return &Story{}
}
