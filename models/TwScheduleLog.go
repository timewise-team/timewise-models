package models

import "time"

type TwScheduleLog struct {
	ID              int             `gorm:"primary_key"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `json:"deleted_at" gorm:"default:null"`
	ScheduleId      int             `json:"schedule_id" gorm:"index"`
	WorkspaceUserId int             `json:"workspace_user_id" gorm:"index"`
	Action          string          `json:"action"`
	FieldChanged    string          `json:"field_changed"`
	OldValue        string          `json:"old_value"`
	NewValue        string          `json:"new_value"`
	Description     string          `json:"description"`
	Schedule        TwSchedule      `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
	WorkspaceUser   TwWorkspaceUser `gorm:"foreignkey:WorkspaceUserId ;association_foreignkey:ID"`
}
