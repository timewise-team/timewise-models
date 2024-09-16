package models
import(

)

type TwScheduleLog struct {
	ID         uint   `gorm:"primary_key"`
	ScheduleId int    `json:"schedule_id" gorm:"index"`
	Log        string `json:"log"`
}
