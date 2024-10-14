package board_columns_dtos

type BoardColumnsRequest struct {
	WorkspaceId int    `json:"workspace_id"`
	Name        string `json:"name"`
	Position    int    `json:"position"`
}
