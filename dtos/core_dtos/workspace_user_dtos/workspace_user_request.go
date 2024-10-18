package workspace_user_dtos

type UpdateWorkspaceUserRoleRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
