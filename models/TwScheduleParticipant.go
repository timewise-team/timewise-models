package models

import (
	"time"
)

type TwScheduleParticipant struct {
	ID               int             `gorm:"primary_key"`
	CreatedAt        *time.Time      `json:"created_at"`
	UpdatedAt        *time.Time      `json:"updated_at"`
	DeletedAt        *time.Time      `json:"deleted_at" gorm:"default:null"`
	ScheduleId       int             `json:"schedule_id" gorm:"index"`
	WorkspaceUserId  int             `json:"workspace_user_id" gorm:"index"`
	Status           string          `json:"status"`
	AssignAt         *time.Time      `json:"assign_at" gorm:"default:null"`
	AssignBy         int             `json:"assign_by"`
	ResponseTime     *time.Time      `json:"response_time" gorm:"default:null"`
	InvitationSentAt *time.Time      `json:"invitation_sent_at" gorm:"default:null"`
	InvitationStatus string          `json:"invitation_status"`
	Schedule         TwSchedule      `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
	WorkspaceUser    TwWorkspaceUser `gorm:"foreignkey:WorkspaceUserId;association_foreignkey:ID"`
	AssignByUser     TwWorkspaceUser `gorm:"foreignkey:AssignBy;association_foreignkey:ID"`
}
