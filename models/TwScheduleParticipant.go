package models

import (
	"time"
)

type TwScheduleParticipant struct {
	ID               int         `gorm:"primary_key"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	DeletedAt        time.Time   `json:"deleted_at" gorm:"default:null"`
	ScheduleId       int         `json:"schedule_id" gorm:"index"`
	UserId           int         `json:"user_id" gorm:"index"`
	Status           string      `json:"status"`
	ResponseTime     time.Time   `json:"response_time"`
	InvitationSentAt time.Time   `json:"invitation_sent_at"`
	Schedule         TwSchedule  `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
	UserEmail        TwUserEmail `gorm:"foreignkey:UserId;association_foreignkey:ID"`
}
