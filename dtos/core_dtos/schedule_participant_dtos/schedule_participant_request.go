package schedule_participant_dtos

import "time"

type InviteToScheduleRequest struct {
	Email      string `json:"email"`
	ScheduleId int    `json:"schedule_id"`
}

type ScheduleParticipantRequest struct {
	Status           *string    `json:"status"`
	AssignAt         *time.Time `json:"assign_at" gorm:"default:null"`
	AssignBy         *int       `json:"assign_by"`
	ResponseTime     *time.Time `json:"response_time" gorm:"default:null"`
	InvitationSentAt *time.Time `json:"invitation_sent_at" gorm:"default:null"`
	InvitationStatus *string    `json:"invitation_status"`
}
