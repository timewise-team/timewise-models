package core_dtos

import "time"

type TwScheduleResponse struct {
	ID                int       `json:"id"`
	WorkspaceID       int       `json:"workspace_id"`
	BoardColumnID     int       `json:"board_column_id"`
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
	AssignedTo        uint64    `json:"assigned_to"`
}

type TwCreateScheduleRequest struct {
	WorkspaceID       *int       `json:"workspace_id,omitempty"`       // Nullable field
	BoardColumnID     *int       `json:"board_column_id,omitempty"`    // Nullable field
	Title             *string    `json:"title,omitempty"`              // Nullable field
	Description       *string    `json:"description,omitempty"`        // Nullable field
	StartTime         *time.Time `json:"start_time,omitempty"`         // Nullable field
	EndTime           *time.Time `json:"end_time,omitempty"`           // Nullable field
	Location          *string    `json:"location,omitempty"`           // Nullable field
	CreatedBy         int        `json:"created_by"`                   // Required field
	Status            *string    `json:"status,omitempty"`             // Nullable field
	AllDay            *bool      `json:"all_day,omitempty"`            // Nullable field (represented by tinyint in SQL)
	Visibility        *string    `json:"visibility,omitempty"`         // Nullable field
	ExtraData         *string    `json:"extra_data,omitempty"`         // Nullable field
	RecurrencePattern *string    `json:"recurrence_pattern,omitempty"` // Nullable field
	AssignedTo        *int       `json:"assigned_to,omitempty"`        // Nullable field
}

type TwCreateShecduleResponse struct {
	ID                int       `json:"id"`
	WorkspaceID       int       `json:"workspace_id"`
	BoardColumnID     int       `json:"board_column_id"`
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
	ExtraData         string    `json:"extra_data,omitempty"`
	IsDeleted         bool      `json:"is_deleted"`
	RecurrencePattern string    `json:"recurrence_pattern"`
	AssignedTo        int       `json:"assigned_to"`
}

type TwUpdateScheduleRequest struct {
	WorkspaceID       *int       `json:"workspace_id,omitempty"`
	BoardColumnID     *int       `json:"board_column_id,omitempty"`
	Title             *string    `json:"title,omitempty"`
	Description       *string    `json:"description,omitempty"`
	StartTime         *time.Time `json:"start_time,omitempty"`
	EndTime           *time.Time `json:"end_time,omitempty"`
	Location          *string    `json:"location,omitempty"`
	Status            *string    `json:"status,omitempty"`
	AllDay            *bool      `json:"all_day,omitempty"`
	Visibility        *string    `json:"visibility,omitempty"`
	ExtraData         *string    `json:"extra_data,omitempty"`
	IsDeleted         *bool      `json:"is_deleted,omitempty"`
	RecurrencePattern *string    `json:"recurrence_pattern,omitempty"`
	AssignedTo        *int       `json:"assigned_to,omitempty"`
}

type TwUpdateScheduleResponse struct {
	ID                int       `json:"id"`
	WorkspaceID       int       `json:"workspace_id"`
	BoardColumnID     int       `json:"board_column_id"`
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
	ExtraData         string    `json:"extra_data,omitempty"`
	IsDeleted         bool      `json:"is_deleted"`
	RecurrencePattern string    `json:"recurrence_pattern"`
	AssignedTo        int       `json:"assigned_to"`
}
