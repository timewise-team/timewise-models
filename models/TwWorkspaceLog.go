package models

import "github.com/jinzhu/gorm"

type TwWorkspaceLog struct {
	gorm.Model
	ID          int    `json:"id"`
	WorkspaceID int    `json:"workspace_id"`
	Log         string `json:"log"`
}
