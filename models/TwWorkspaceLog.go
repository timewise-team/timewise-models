package models

import "time"

type TwWorkspaceLog struct {
	ID          int       `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	WorkspaceId int       `json:"workspace_id" gorm:"index"`
	Log         string    `json:"log"`
}
