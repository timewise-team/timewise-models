package models

import "github.com/jinzhu/gorm"

type TwScheduleLog struct {
	gorm.Model
	ScheduleId int    `json:"schedule_id" gorm:"index"`
	Log        string `json:"log"`
}
