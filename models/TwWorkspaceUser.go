package models

import "time"

type TwWorkspaceUser struct {
	ID           int       `gorm:"primary_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	UserId       int       `json:"user_id"`
	WorkspaceId  int       `json:"workspace_id" gorm:"index"`
	WorkspaceKey string    `json:"workspace_key"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	IsActive     bool      `json:"is_active"`
	ExtraData    string    `json:"extra_data"`
}
