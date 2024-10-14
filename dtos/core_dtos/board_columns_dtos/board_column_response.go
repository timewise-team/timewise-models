package board_columns_dtos

import (
	dtos "github.com/timewise-team/timewise-models/dtos/core_dtos"
	"time"
)

type BoardColumnsResponse struct {
	ID          int                                        `json:"id"`
	CreatedAt   time.Time                                  `json:"created_at"`
	UpdatedAt   time.Time                                  `json:"updated_at"`
	DeletedAt   time.Time                                  `json:"deleted_at"`
	WorkspaceId int                                        `json:"workspace_id"`
	Name        string                                     `json:"name"`
	Position    int                                        `json:"position"`
	Schedules   []dtos.TwScheduleListInBoardColumnResponse `json:"schedules"`
}
