package models


type TwScheduleLog struct {
	ID         int    `json:"id"`
	ScheduleID int    `json:"schedule_id"`
	Log        string `json:"log"`
}
