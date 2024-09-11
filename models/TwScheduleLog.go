package models

import "github.com/jinzhu/gorm"

type TwScheduleLog struct {
	gorm.Model
	ID         int    `json:"id"`
	ScheduleID int    `json:"schedule_id"`
	Log        string `json:"log"`
}
