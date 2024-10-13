package workspace_user_dtos

import "time"

type GetWorkspaceUserListResponse struct {
	ID             int       `json:"id"`
	UserEmailId    int       `json:"user_email_id"`
	WorkspaceId    int       `json:"workspace_id"`
	WorkspaceKey   string    `json:"workspace_key"`
	Role           string    `json:"role"`
	Status         string    `json:"status"`
	IsActive       bool      `json:"is_active"`
	IsVerified     bool      `json:"is_verified"`
	ExtraData      string    `json:"extra_data"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	Email          string    `json:"email"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	ProfilePicture string    `json:"profile_picture"`
}
