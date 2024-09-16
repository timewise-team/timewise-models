package models

import "github.com/jinzhu/gorm"

type TwWorkspaceLog struct {
	gorm.Model
	WorkspaceId int    `json:"workspace_id" gorm:"index"`
	Log         string `json:"log"`
}