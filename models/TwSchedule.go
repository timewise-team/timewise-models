package models

import (
	"time"
)

type TwSchedule struct {
	ID                int       `gorm:"primary_key"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
	WorkspaceId       int       `json:"workspace_id" gorm:"index"`
	BoardColumnId     int       `json:"board_column_id" gorm:"index"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	Location          string    `json:"location"`
	CreatedBy         int       `json:"created_by"`
	Status            string    `json:"status"`
	AllDay            bool      `json:"all_day"`
	Visibility        string    `json:"visibility"`
	ExtraData         string    `json:"extra_data"`
	IsDeleted         bool      `json:"is_deleted"`
	RecurrencePattern string    `json:"recurrence_pattern"`
	AssignedTo        int       `json:"assigned_to"`
}
