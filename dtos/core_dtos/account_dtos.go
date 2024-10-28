package core_dtos

import "time"

type GetUserResponseDto struct {
	ID                   int       `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeteledAt            time.Time `json:"deleted_at"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Email                []string  `json:"email"`
	ProfilePicture       string    `json:"profile_picture"`
	Timezone             string    `json:"timezone"`
	Locale               string    `json:"locale"`
	GoogleId             string    `json:"google_id"`
	IsVerified           bool      `json:"is_verified"`
	IsActive             bool      `json:"is_active"`
	LastLoginAt          time.Time `json:"last_login_at"`
	NotificationSettings string    `json:"notification_settings"`
	CalendarSettings     string    `json:"calendar_settings"`
	Role                 string    `json:"role"`
}

type UpdateProfileRequestDto struct {
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	ProfilePicture       string `json:"profile_picture"`
	NotificationSettings string `json:"notification_settings"`
	CalendarSettings     string `json:"calendar_settings"`
}

type ChangePasswordRequestDto struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
