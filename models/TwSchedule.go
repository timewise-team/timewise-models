package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwSchedule struct {
	gorm.Model
	ID                int       `json:"id"`
	WorkspaceID       int       `json:"workspace_id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
	Location          string    `json:"location"`
	CreatedBy         int       `json:"created_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Status            string    `json:"status"`
	AllDay            bool      `json:"all_day"`
	Visibility        string    `json:"visibility"`
	ExtraData         string    `json:"extra_data"`
	IsDeleted         bool      `json:"is_deleted"`
	RecurrencePattern string    `json:"recurrence_pattern"`
	AssignedTo        int       `json:"assigned_to"`
}
