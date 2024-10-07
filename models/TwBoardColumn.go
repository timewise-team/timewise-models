package models

import "time"

type TwBoardColumn struct {
	ID          int         `gorm:"primary_key"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	DeletedAt   time.Time   `json:"deleted_at" gorm:"default:null"`
	WorkspaceId int         `json:"workspace_id" gorm:"index"`
	Name        string      `json:"name" gorm:"varchar(255)"`
	Position    int         `json:"position" gorm:"index"`
	Workspace   TwWorkspace `gorm:"foreignkey:WorkspaceId;association_foreignkey:ID"`
}
