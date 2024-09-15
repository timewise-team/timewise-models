package models

import (
	"time"
)

type TwUser struct {
	ID                   int       `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Email                string    `json:"email"`
	PasswordHash         string    `json:"password_hash"`
	ProfilePicture       string    `json:"profile_picture"`
	Timezone             string    `json:"timezone"`
	Locale               string    `json:"locale"`
	IsVerified           bool      `json:"is_verified"`
	IsActive             bool      `json:"is_active"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	LastLoginAt          time.Time `json:"last_login_at"`
	NotificationSettings string    `json:"notification_settings"`
	CalendarSettings     string    `json:"calendar_settings"`
	Role                 string    `json:"role"`
}

func (TwUser) TableName() string {
	return "tw_user"
}
