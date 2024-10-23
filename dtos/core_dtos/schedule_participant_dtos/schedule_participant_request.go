package schedule_participant_dtos

type InviteToScheduleRequest struct {
	Email      string `json:"email"`
	ScheduleId int    `json:"schedule_id"`
}
