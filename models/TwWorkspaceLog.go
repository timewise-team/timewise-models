package models


type TwWorkspaceLog struct {
	ID          int    `json:"id"`
	WorkspaceID int    `json:"workspace_id"`
	Log         string `json:"log"`
}
