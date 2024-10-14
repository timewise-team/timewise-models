package models

import (
	"time"
)

type TwSchedule struct {
	ID                int             `gorm:"primary_key"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         time.Time       `json:"deleted_at" gorm:"default:null"`
	WorkspaceId       int             `json:"workspace_id" gorm:"index"`
	BoardColumnId     int             `json:"board_column_id" gorm:"index"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	StartTime         time.Time       `json:"start_time" gorm:"default:null"`
	EndTime           time.Time       `json:"end_time" gorm:"default:null"`
	Location          string          `json:"location"`
	CreatedBy         int             `json:"workspace_user_id" gorm:"index"`
	Status            string          `json:"status"`
	AllDay            bool            `json:"all_day"`
	Visibility        string          `json:"visibility"`
	VideoTranscript   string          `json:"video_transcript"`
	ExtraData         string          `json:"extra_data"`
	IsDeleted         bool            `json:"is_deleted"`
	RecurrencePattern string          `json:"recurrence_pattern"`
	Workspace         TwWorkspace     `gorm:"foreignkey:WorkspaceId;association_foreignkey:ID"`
	BoardColumn       TwBoardColumn   `gorm:"foreignkey:BoardColumnId;association_foreignkey:ID"`
	WorkspaceUser     TwWorkspaceUser `gorm:"foreignkey:CreatedBy;association_foreignkey:ID"`
}
