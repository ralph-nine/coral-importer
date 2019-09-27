//go:generate easyjson -all users.go
package coral

type UserProfile struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	DastIssuedAt Time   `json:"lastIssuedAt"`
}

type UserNotifications struct {
	OnReply         bool   `json:"onReply"`
	OnFeatured      bool   `json:"onFeatured"`
	OnStaffReplies  bool   `json:"onStaffReplies"`
	OnModeration    bool   `json:"onModeration"`
	DigestFrequency string `json:"digestFrequency" validate:"oneof=NONE DAILY HOURLY"`
}

func NewUserNotifications() UserNotifications {
	return UserNotifications{
		DigestFrequency: "NONE",
	}
}

type UserSuspensionStatus struct {
	History []string `json:"history"`
}

type UserBanStatus struct {
	Active  bool     `json:"active"`
	History []string `json:"history"`
}

type UserUsernameStatusHistory struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedBy string `json:"createdBy"`
	CreatedAt Time   `json:"createdAt"`
}

type UserUsernameStatus struct {
	History []UserUsernameStatusHistory `json:"history"`
}

func NewUserUsernameStatus() UserUsernameStatus {
	return UserUsernameStatus{
		History: []UserUsernameStatusHistory{},
	}
}

type UserPremodStatus struct {
	Active  bool     `json:"active"`
	History []string `json:"history"`
}

type UserStatus struct {
	SuspensionStatus UserSuspensionStatus `json:"suspension"`
	BanStatus        UserBanStatus        `json:"ban"`
	UsernameStatus   UserUsernameStatus   `json:"username"`
	PremodStatus     UserPremodStatus     `json:"premod"`
}

func NewUserStatus() UserStatus {
	return UserStatus{
		SuspensionStatus: UserSuspensionStatus{
			History: []string{},
		},
		BanStatus: UserBanStatus{
			History: []string{},
		},
		UsernameStatus: NewUserUsernameStatus(),
		PremodStatus: UserPremodStatus{
			History: []string{},
		},
	}
}

type User struct {
	TenantID      string            `json:"tenantID" validate:"required"`
	ID            string            `json:"id" conform:"trim" validate:"required"`
	Username      string            `json:"username" validate:"required"`
	Email         string            `json:"email" validate:"email"`
	Profiles      []UserProfile     `json:"profiles,omitempty"`
	Role          string            `json:"role" validate:"required,oneof=COMMENTER STAFF MODERATOR ADMIN"`
	Notifications UserNotifications `json:"notifications"`
	Status        UserStatus        `json:"status"`
	CreatedAt     Time              `json:"createdAt" validate:"required"`
	Imported      bool              `json:"imported"`
}

func NewUser(tenantID string) *User {
	return &User{
		TenantID:      tenantID,
		Notifications: NewUserNotifications(),
		Status:        NewUserStatus(),
		Profiles:      []UserProfile{},
		Role:          "COMMENTER",
		Imported:      true,
	}
}
