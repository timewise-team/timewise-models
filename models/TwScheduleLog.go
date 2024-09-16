package models



type TwScheduleLog struct {
	ID         int    `gorm:"primary_key"`			
	ScheduleId int    `json:"schedule_id" gorm:"index"`
	Log        string `json:"log"`
}
