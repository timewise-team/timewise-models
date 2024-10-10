package create_workspace_dtos

type CreateWorkspaceRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Key         string `json:"key"`
	Type        string `json:"type"`
	IsDeleted   bool   `json:"is_deleted"`
	Email       string `json:"email"`
}
