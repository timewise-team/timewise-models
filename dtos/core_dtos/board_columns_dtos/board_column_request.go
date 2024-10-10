package board_columns_dtos

type CreateBoardColumnRequest struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}
