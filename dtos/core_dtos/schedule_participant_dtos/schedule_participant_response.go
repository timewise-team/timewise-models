package schedule_participant_dtos

import "time"

type ScheduleParticipantInfo struct {
	Id                  int    `json:"id"`
	ScheduleId          int    `json:"schedule_id"`
	WorkspaceUserId     int    `json:"workspace_user_id"`
	Status              string `json:"status"`
	AssignAt            string `json:"assign_at"`
	AssignBy            int    `json:"assign_by"`
	ResponseTime        string `json:"response_time"`
	InvitationSentAt    string `json:"invitation_sent_at"`
	InvitationStatus    string `json:"invitation_status"`
	Role                string `json:"role"`
	StatusWorkspaceUser string `json:"status_workspace_user"`
	IsVerified          bool   `json:"is_verified"`
	UserId              int    `json:"user_id" `
	Email               string `json:"email"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	ProfilePicture      string `json:"profile_picture"`
}

type ScheduleParticipantResponse struct {
	ID               int        `gorm:"primary_key"`
	ScheduleId       int        `json:"schedule_id" gorm:"index"`
	WorkspaceUserId  int        `json:"workspace_user_id" gorm:"index"`
	Status           string     `json:"status"`
	AssignAt         *time.Time `json:"assign_at" gorm:"default:null"`
	AssignBy         int        `json:"assign_by"`
	ResponseTime     *time.Time `json:"response_time" gorm:"default:null"`
	InvitationSentAt *time.Time `json:"invitation_sent_at" gorm:"default:null"`
	InvitationStatus string     `json:"invitation_status"`
}
