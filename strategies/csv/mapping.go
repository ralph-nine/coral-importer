package csv

import (
	"github.com/pkg/errors"
)

func ValidateRequired(fields []string, required int) error {
	if len(fields) < required {
		return errors.Errorf("row had %d fields but %d are required", len(fields), required)
	}

	for i := 0; i < required-1; i++ {
		value := fields[i]
		if value == "" {
			return errors.Errorf("field %d was required, but was empty", i)
		}
	}

	return nil
}

const CommentColumns = 7

type Comment struct {
	ID        string
	AuthorID  string
	StoryID   string
	CreatedAt string
	Body      string
	ParentID  string
	Status    string
}

func ParseComment(fields []string) (*Comment, error) {
	if err := ValidateRequired(fields, 5); err != nil {
		return nil, err
	}

	return &Comment{
		ID:        fields[0],
		AuthorID:  fields[1],
		StoryID:   fields[2],
		CreatedAt: fields[3],
		Body:      fields[4],
		ParentID:  fields[5],
		Status:    fields[6],
	}, nil
}

const StoryColumns = 6

type Story struct {
	ID          string
	URL         string
	Title       string
	Author      string
	PublishedAt string
	ClosedAt    string
}

func ParseStory(fields []string) (*Story, error) {
	if err := ValidateRequired(fields, 2); err != nil {
		return nil, err
	}

	return &Story{
		ID:          fields[0],
		URL:         fields[1],
		Title:       fields[2],
		Author:      fields[3],
		PublishedAt: fields[4],
		ClosedAt:    fields[5],
	}, nil
}

const UserColumns = 6

type User struct {
	ID        string
	Email     string
	Username  string
	Role      string
	Banned    string
	CreatedAt string
}

func ParseUser(fields []string) (*User, error) {
	if err := ValidateRequired(fields, 3); err != nil {
		return nil, err
	}

	return &User{
		ID:        fields[0],
		Email:     fields[1],
		Username:  fields[2],
		Role:      fields[3],
		Banned:    fields[4],
		CreatedAt: fields[5],
	}, nil
}
