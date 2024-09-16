package models

type TwWorkspaceLog struct {
	ID          uint   `gorm:"primary_key"`
	WorkspaceId int    `json:"workspace_id" gorm:"index"`
	Log         string `json:"log"`
}
