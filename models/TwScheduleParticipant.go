package models

import (
	"time"
)

type TwScheduleParticipant struct {
	ID               int       `json:"id"`
	ScheduleID       int       `json:"schedule_id"`
	UserID           int       `json:"user_id"`
	Status           string    `json:"status"`
	ResponseTime     time.Time `json:"response_time"`
	Role             string    `json:"role"`
	InvitationSentAt time.Time `json:"invitation_sent_at"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
