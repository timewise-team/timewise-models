package models

import (
	"github.com/jinzhu/gorm"
)

type TwBoardColumn struct {
	gorm.Model
	WorkspaceId int    `json:"workspace_id" gorm:"index"`
	Name        string `json:"name"`
}
