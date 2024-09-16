package models

type TwWorkspaceUser struct {
	ID           uint   `gorm:"primary_key"`
	UserId       int    `json:"user_id"`
	WorkspaceId  int    `json:"workspace_id" gorm:"index"`
	WorkspaceKey string `json:"workspace_key"`
	Status       string `json:"status"`
	IsActive     bool   `json:"is_active"`
	ExtraData    string `json:"extra_data"`
}
