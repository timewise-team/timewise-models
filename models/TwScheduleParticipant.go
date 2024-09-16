package models

import (
	"time"
)

type TwScheduleParticipant struct {
	ID               int       `gorm:"primary_key"`
	ScheduleId       int       `json:"schedule_id" gorm:"index"`
	UserId           int       `json:"user_id" gorm:"index"`
	Status           string    `json:"status"`
	ResponseTime     time.Time `json:"response_time"`
	Role             string    `json:"role"`
	InvitationSentAt time.Time `json:"invitation_sent_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
