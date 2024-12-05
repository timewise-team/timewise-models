package core_dtos

import (
	"github.com/timewise-team/timewise-models/dtos/core_dtos/schedule_participant_dtos"
	"time"
)

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
	VideoTranscript   *string   `json:"video_transcript"`
	ExtraData         string    `json:"extra_data"`
	IsDeleted         bool      `json:"is_deleted"`
	RecurrencePattern string    `json:"recurrence_pattern"`
	Position          int       `json:"position"`
	Priority          string    `json:"priority"`
	//AssignedTo        []int     `json:"assigned_to"`
}

type TwGetScheduleResponse struct {
	ID                int    `json:"id"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	DeletedAt         string `json:"deleted_at" gorm:"default:null"`
	WorkspaceId       int    `json:"workspace_id" gorm:"index"`
	BoardColumnId     int    `json:"board_column_id" gorm:"index"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	StartTime         string `json:"start_time" gorm:"default:null"`
	EndTime           string `json:"end_time" gorm:"default:null"`
	Location          string `json:"location"`
	CreatedBy         int    `json:"workspace_user_id" gorm:"index"`
	Status            string `json:"status"`
	AllDay            bool   `json:"all_day"`
	Visibility        string `json:"visibility"`
	VideoTranscript   string `json:"video_transcript"`
	ExtraData         string `json:"extra_data"`
	IsDeleted         bool   `json:"is_deleted"`
	RecurrencePattern string `json:"recurrence_pattern"`
	Position          int    `json:"position"`
	Priority          string `json:"priority"`
}

type TwCreateScheduleRequest struct {
	WorkspaceID     *int    `json:"workspace_id,omitempty"` // Nullable field
	WorkspaceUserID *int    `json:"workspace_user_id,omitempty"`
	BoardColumnID   *int    `json:"board_column_id,omitempty"` // Nullable field
	Title           *string `json:"title,omitempty"`           // Nullable field
	Description     *string `json:"description,omitempty"`     // Nullable field
	StartTime       *string `json:"start_time,omitempty"`
	EndTime         *string `json:"end_time,omitempty"`
}

type TwCreateShecduleResponse struct {
	ID            int       `json:"id"`
	WorkspaceID   int       `json:"workspace_id"`
	BoardColumnID int       `json:"board_column_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Position      int       `json:"position"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
}

type TwUpdateScheduleRequest struct {
	WorkspaceID *int `json:"workspace_id,omitempty"`
	//BoardColumnID     *int       `json:"board_column_id,omitempty"`
	Title             *string `json:"title,omitempty"`
	Description       *string `json:"description,omitempty"`
	StartTime         *string `json:"start_time,omitempty"`
	EndTime           *string `json:"end_time,omitempty"`
	Location          *string `json:"location,omitempty"`
	Status            *string `json:"status,omitempty"`
	AllDay            *bool   `json:"all_day,omitempty"`
	Visibility        *string `json:"visibility,omitempty"`
	VideoTranscript   *string `json:"video_transcript"`
	ExtraData         *string `json:"extra_data,omitempty"`
	RecurrencePattern *string `json:"recurrence_pattern,omitempty"`
	//Position          *int       `json:"position,omitempty"`
	Priority *string `json:"priority,omitempty"`
}

type TwUpdateSchedulePosition struct {
	BoardColumnID *int `json:"board_column_id,omitempty"`
	Position      *int `json:"position,omitempty"`
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
	VideoTranscript   string    `json:"video_transcript"`
	ExtraData         string    `json:"extra_data,omitempty"`
	IsDeleted         bool      `json:"is_deleted"`
	RecurrencePattern string    `json:"recurrence_pattern"`
	Position          int       `json:"position"`
	Priority          string    `json:"priority"`
}

type TwUpdateScheduleTimeResponse struct {
	ID                int    `json:"id"`
	WorkspaceID       int    `json:"workspace_id"`
	BoardColumnID     int    `json:"board_column_id"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	Location          string `json:"location"`
	CreatedBy         int    `json:"created_by"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Status            string `json:"status"`
	AllDay            bool   `json:"all_day"`
	Visibility        string `json:"visibility"`
	VideoTranscript   string `json:"video_transcript"`
	ExtraData         string `json:"extra_data,omitempty"`
	IsDeleted         bool   `json:"is_deleted"`
	RecurrencePattern string `json:"recurrence_pattern"`
	Position          int    `json:"position"`
	Priority          string `json:"priority"`
}

type TwScheduleListInBoardColumnResponse struct {
	ID                   int                                                 `json:"id"`
	WorkspaceID          int                                                 `json:"workspace_id"`
	BoardColumnID        int                                                 `json:"board_column_id"`
	Title                string                                              `json:"title"`
	Description          string                                              `json:"description"`
	StartTime            time.Time                                           `json:"start_time"`
	EndTime              time.Time                                           `json:"end_time"`
	Location             string                                              `json:"location"`
	CreatedBy            int                                                 `json:"created_by"`
	CreatedAt            time.Time                                           `json:"created_at"`
	UpdatedAt            time.Time                                           `json:"updated_at"`
	Status               string                                              `json:"status"`
	AllDay               bool                                                `json:"all_day"`
	Visibility           string                                              `json:"visibility"`
	VideoTranscript      string                                              `json:"video_transcript"`
	ExtraData            string                                              `json:"extra_data,omitempty"`
	IsDeleted            bool                                                `json:"is_deleted"`
	RecurrencePattern    string                                              `json:"recurrence_pattern"`
	Position             int                                                 `json:"position"`
	Priority             string                                              `json:"priority"`
	ScheduleParticipants []schedule_participant_dtos.ScheduleParticipantInfo `json:"schedule_participants"`
	Documents            int                                                 `json:"documents_count"`
	Comments             int                                                 `json:"comments_count"`
}
