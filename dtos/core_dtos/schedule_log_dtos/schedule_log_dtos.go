package schedule_log_dtos

import "time"

type TwScheduleLogResponse struct {
	ID                  int       `gorm:"primary_key"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	ScheduleId          int       `json:"schedule_id" gorm:"index"`
	WorkspaceUserId     int       `json:"workspace_user_id" gorm:"index"`
	Action              string    `json:"action"`
	FieldChanged        string    `json:"field_changed"`
	OldValue            string    `json:"old_value"`
	NewValue            string    `json:"new_value"`
	Description         string    `json:"description"`
	Role                string    `json:"role"`
	StatusWorkspaceUser string    `json:"status_workspace_user"`
	IsVerified          bool      `json:"is_verified"`
	UserId              int       `json:"user_id" `
	Email               string    `json:"email"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	ProfilePicture      string    `json:"profile_picture"`
}
