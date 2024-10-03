package models

import (
	"time"
)

type TwReminder struct {
	ID              int             `gorm:"primary_key"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `json:"deleted_at" gorm:"default:null"`
	ScheduleId      int             `json:"schedule_id" gorm:"index"`
	ReminderTime    time.Time       `json:"reminder_time"`
	Method          string          `json:"method"`
	Type            string          `json:"type"`
	IsSent          bool            `json:"is_sent"`
	WorkspaceUserID int             `json:"workspace_user_id" gorm:"index"`
	WorkspaceUser   TwWorkspaceUser `gorm:"foreignkey:WorkspaceUserID;association_foreignkey:ID"`
	Schedule        TwSchedule      `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
}
