//go:generate easyjson -all models.go
package legacy

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/coralproject/coral-importer/common/coral"
)

// Action is the Action as exported from MongoDB from legacy Talk.
type Action struct {
	ID         string                 `json:"id"`
	ActionType string                 `json:"action_type"`
	GroupID    string                 `json:"group_id"`
	ItemID     string                 `json:"item_id"`
	ItemType   string                 `json:"item_type"`
	UserID     *string                `json:"user_id"`
	Metadata   map[string]interface{} `json:"metadata"`
	CreatedAt  coral.Time             `json:"created_at"`
}

func TranslateCommentAction(tenantID string, action *Action) *coral.CommentAction {
	commentAction := coral.NewCommentAction(tenantID)
	commentAction.ID = action.ID
	if action.ActionType == "FLAG" {
		commentAction.ActionType = "FLAG"
		switch action.GroupID {
		case "BANNED_WORD":
			commentAction.Reason = "COMMENT_DETECTED_BANNED_WORD"
		case "COMMENT_OFFENSIVE":
			commentAction.Reason = "COMMENT_REPORTED_OFFENSIVE"
		case "COMMENT_OTHER":
			commentAction.Reason = "COMMENT_REPORTED_OTHER"
		case "COMMENT_SPAM":
			commentAction.Reason = "COMMENT_REPORTED_SPAM"
		case "LINKS":
			commentAction.Reason = "COMMENT_DETECTED_LINKS"
		case "SPAM_COMMENT":
			commentAction.Reason = "COMMENT_DETECTED_SPAM"
		case "TRUST":
			commentAction.Reason = "COMMENT_DETECTED_RECENT_HISTORY"
		case "TOXIC_COMMENT":
			commentAction.Reason = "COMMENT_DETECTED_TOXIC"
		case "":
		default:
		}

		if action.Metadata != nil {
			message, ok := action.Metadata["message"].(string)
			if ok {
				commentAction.AdditionalDetails = message
			}
		}
	} else if action.ActionType == "DONTAGREE" {
		commentAction.ActionType = "DONT_AGREE"
	} else {
		commentAction.ActionType = "REACTION"
	}
	commentAction.CommentID = action.ItemID
	commentAction.UserID = action.UserID
	commentAction.CreatedAt.Time = action.CreatedAt.Time

	// v4 did not have revision ID's, so use the comment ID which will always be
	// used as the most recent body history item anyways when we import the
	// comments.
	commentAction.CommentRevisionID = action.ItemID

	// The following must be processed when we've loaded all the comments in via
	// a second pass.
	commentAction.StoryID = ""

	return commentAction
}

type CommentBodyHistory struct {
	Body      string     `json:"body"`
	CreatedAt coral.Time `json:"created_at"`
}

type CommentTag struct {
	AssignedBy *string `json:"assigned_by"`
	Tag        struct {
		Name string `json:"name"`
	} `json:"tag"`
	CreatedAt coral.Time `json:"created_at"`
}

type Comment struct {
	ID          string               `json:"id"`
	AssetID     string               `json:"asset_id"`
	Status      string               `json:"status"`
	BodyHistory []CommentBodyHistory `json:"body_history"`
	Tags        []CommentTag         `json:"tags"`
	ParentID    *string              `json:"parent_id"`
	AuthorID    string               `json:"author_id"`
	DeletedAt   *coral.Time          `json:"deleted_at"`
	CreatedAt   coral.Time           `json:"created_at"`
	UpdatedAt   coral.Time           `json:"updated_at"`
}

func TranslateCommentStatus(status string) string {
	if status == "ACCEPTED" {
		return "APPROVED"
	}
	return status
}

func TranslateComment(tenantID string, in *Comment) *coral.Comment {
	comment := coral.NewComment(tenantID)
	comment.ID = in.ID
	if in.ParentID != nil {
		comment.ParentID = *in.ParentID
		comment.ParentRevisionID = *in.ParentID
	}
	comment.AuthorID = in.AuthorID
	// comment.Tags
	comment.CreatedAt.Time = in.CreatedAt.Time
	comment.StoryID = in.AssetID
	comment.Status = TranslateCommentStatus(in.Status)
	for _, tag := range in.Tags {
		commentTag := coral.CommentTag{
			Type:      tag.Tag.Name,
			CreatedAt: tag.CreatedAt,
		}

		if tag.AssignedBy != nil {
			commentTag.CreatedBy = *tag.AssignedBy
		}

		comment.Tags = append(comment.Tags, commentTag)
	}
	if in.DeletedAt == nil {
		revisionLength := len(in.BodyHistory)

		if revisionLength <= 0 {
			panic(fmt.Sprintf("%s with deletedAt: %v which is %v", in.ID, in.DeletedAt, in.DeletedAt == nil))
		}

		comment.Revisions = make([]coral.Revision, revisionLength)
		for i, revision := range in.BodyHistory {
			comment.Revisions[i] = coral.Revision{
				ID:           comment.ID + "-" + fmt.Sprintf("%d", i),
				Body:         coral.HTML(revision.Body),
				Metadata:     coral.RevisionMetadata{},
				ActionCounts: map[string]int{},
			}
			comment.Revisions[i].CreatedAt.Time = revision.CreatedAt.Time
		}

		// Ensure that the last revision ID is the comment's ID.
		comment.Revisions[revisionLength-1].ID = comment.ID
	} else {
		comment.DeletedAt = in.DeletedAt
	}

	return comment
}

type Asset struct {
	ID            string      `json:"id"`
	URL           string      `json:"url"`
	ClosedAt      *coral.Time `json:"closedAt"`
	ClosedMessage *string     `json:"closedMessage"`
	CreatedAt     coral.Time  `json:"created_at"`
	Scraped       *coral.Time `json:"scraped"`
	Metadata      interface{} `json:"metadata"`
	Settings      interface{} `json:"settings"`
	// Tags

	// Scraped.
	Title           *string     `json:"title"`
	Author          *string     `json:"author"`
	Description     *string     `json:"description"`
	Image           *string     `json:"image"`
	Section         *string     `json:"section"`
	ModifiedDate    *coral.Time `json:"modified_date"`
	PublicationDate *coral.Time `json:"publication_date"`
}

func TranslateAsset(tenantID string, asset *Asset) *coral.Story {
	story := coral.NewStory(tenantID)
	story.ID = asset.ID
	story.URL = asset.URL
	if asset.Title != nil {
		story.Metadata.Title = *asset.Title
	}
	if asset.Author != nil {
		story.Metadata.Author = *asset.Author
	}
	if asset.Description != nil {
		story.Metadata.Description = *asset.Description
	}
	if asset.Image != nil {
		story.Metadata.Image = *asset.Image
	}
	if asset.Section != nil {
		story.Metadata.Section = *asset.Section
	}
	if asset.PublicationDate != nil {
		story.Metadata.PublishedAt = asset.PublicationDate
	}

	return story
}

type UserProfile struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

type UserToken struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type UserNotificationSettings struct {
	OnFeatured      *bool       `json:"onFeatured"`
	OnModeration    *bool       `json:"onModeration"`
	OnReply         *bool       `json:"onReply"`
	OnStaffReply    *bool       `json:"onStaffReply"`
	DigestFrequency interface{} `json:"digestFrequency"`
}

type UserNotifications struct {
	Settings *UserNotificationSettings `json:"settings"`
}

type UserMetadata struct {
	Notifications *UserNotifications `json:"notifications"`
}

type User struct {
	ID           string        `json:"id"`
	Username     string        `json:"username"`
	Role         string        `json:"role"`
	Password     string        `json:"password"`
	IgnoredUsers []string      `json:"ignoresUsers"`
	Profiles     []UserProfile `json:"profiles"`
	Tokens       []UserToken   `json:"tokens"`
	CreatedAt    coral.Time    `json:"created_at"`
	Metadata     *UserMetadata `json:"metadata"`
}

func TranslateUser(tenantID string, in *User) *coral.User {
	user := coral.NewUser(tenantID)
	user.ID = in.ID
	user.Username = in.Username
	user.Role = in.Role
	user.CreatedAt = in.CreatedAt
	user.IgnoredUsers = make([]coral.IgnoredUser, len(in.IgnoredUsers))
	for i, ignoredUserID := range in.IgnoredUsers {
		user.IgnoredUsers[i] = coral.IgnoredUser{
			ID:        ignoredUserID,
			CreatedAt: user.CreatedAt,
		}
	}
	// TODO: status.suspension
	// TODO: status.banned
	user.Profiles = make([]coral.UserProfile, len(in.Profiles))
	for i, profile := range in.Profiles {
		switch profile.Provider {
		case "local":
			user.Profiles[i] = coral.UserProfile{
				ID:         profile.ID,
				Type:       "local",
				Password:   in.Password,
				PasswordID: in.ID,
			}
			user.Email = profile.ID
			break
		case "facebook":
			user.Profiles[i] = coral.UserProfile{
				ID:   profile.ID,
				Type: "facebook",
			}
			break
		case "google":
			user.Profiles[i] = coral.UserProfile{
				ID:   profile.ID,
				Type: "google",
			}
			break
		default:
			panic(errors.Errorf("unsupported profile provider: %s: %v", profile.Provider, in.ID))
		}
	}
	for _, token := range in.Tokens {
		if token.Active {
			user.Tokens = append(user.Tokens, coral.UserToken{
				ID:        token.ID,
				Name:      token.Name,
				CreatedAt: user.CreatedAt,
			})
		}
	}
	if in.Metadata != nil && in.Metadata.Notifications != nil && in.Metadata.Notifications.Settings != nil {
		if in.Metadata.Notifications.Settings.OnReply != nil {
			user.Notifications.OnReply = *in.Metadata.Notifications.Settings.OnReply
		}
		if in.Metadata.Notifications.Settings.OnFeatured != nil {
			user.Notifications.OnFeatured = *in.Metadata.Notifications.Settings.OnFeatured
		}
		if in.Metadata.Notifications.Settings.OnStaffReply != nil {
			user.Notifications.OnStaffReplies = *in.Metadata.Notifications.Settings.OnStaffReply
		}
		if in.Metadata.Notifications.Settings.OnModeration != nil {
			user.Notifications.OnModeration = *in.Metadata.Notifications.Settings.OnModeration
		}
		if in.Metadata.Notifications.Settings.DigestFrequency != nil {
			// Sometimes it seems the digestFrequency is `false` instead of a
			// string, this is a mitigation for that.
			if digestFrequency, ok := in.Metadata.Notifications.Settings.DigestFrequency.(string); ok {
				user.Notifications.DigestFrequency = digestFrequency
			} else {
				user.Notifications.DigestFrequency = "NONE"
			}
		} else {
			user.Notifications.DigestFrequency = "NONE"
		}
	}

	return user
}
