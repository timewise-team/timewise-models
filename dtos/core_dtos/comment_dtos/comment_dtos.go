package comment_dtos

import "time"

type TwCommentResponse struct {
	ID                  int       `gorm:"primary_key"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ScheduleId          int       `json:"schedule_id"`
	WorkspaceUserId     int       `json:"workspace_user_id"`
	Commenter           string    `json:"commenter"`
	Content             string    `json:"content"`
	IsDeleted           bool      `json:"is_deleted"`
	Role                string    `json:"role"`
	StatusWorkspaceUser string    `json:"status_workspace_user"`
	IsVerified          bool      `json:"is_verified"`
	UserId              int       `json:"user_id" `
	Email               string    `json:"email"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	ProfilePicture      string    `json:"profile_picture"`
}

type CommentRequestDTO struct {
	ScheduleId *int    `json:"schedule_id"`
	Commenter  *string `json:"commenter"`
	Content    *string `json:"content"`
}

type CommentResponseDTO struct {
	ID              int        `gorm:"primary_key"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	ScheduleId      int        `json:"schedule_id"`
	WorkspaceUserId int        `json:"workspace_user_id"`
	Commenter       string     `json:"commenter"`
	Content         string     `json:"content"`
	IsDeleted       bool       `json:"is_deleted"`
}
