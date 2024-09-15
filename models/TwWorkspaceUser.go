package models

type TwWorkspaceUser struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	WorkspaceID  int    `json:"workspace_id"`
	WorkspaceKey string `json:"workspace_key"`
	Status       string `json:"status"`
	IsActive     bool   `json:"is_active"`
	ExtraData    string `json:"extra_data"`
}
