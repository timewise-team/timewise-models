package models

import "time"

type TwWorkspaceLog struct {
	ID           int         `gorm:"primary_key"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	DeletedAt    time.Time   `json:"deleted_at" gorm:"default:null"`
	WorkspaceId  int         `json:"workspace_id" gorm:"index"`
	UserId       int         `json:"user_id" gorm:"index"`
	Action       string      `json:"action"`
	FieldChanged string      `json:"field_changed"`
	OldValue     string      `json:"old_value"`
	NewValue     string      `json:"new_value"`
	Description  string      `json:"description"`
	Workspace    TwWorkspace `gorm:"foreignkey:WorkspaceId;association_foreignkey:ID"`
	UserEmail    TwUserEmail `gorm:"foreignkey:UserId;association_foreignkey:ID"`
}
