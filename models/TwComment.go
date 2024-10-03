package models

import (
	"time"
)

type TwComment struct {
	ID              int             `gorm:"primary_key"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       time.Time       `json:"deleted_at" gorm:"default:null"`
	ScheduleId      int             `json:"schedule_id"`
	WorkspaceUserId int             `json:"workspace_user_id"`
	Commenter       string          `json:"commenter"`
	Content         string          `json:"content"`
	IsDeleted       bool            `json:"is_deleted"`
	Schedule        TwSchedule      `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
	WorkspaceUser   TwWorkspaceUser `gorm:"foreignkey:WorkspaceUserId;association_foreignkey:ID"`
}
