package create_workspace_dtos

type CreateWorkspaceRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Email       string `json:"email"`
}
