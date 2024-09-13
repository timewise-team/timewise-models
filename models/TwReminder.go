package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwReminder struct {
	gorm.Model
	ScheduleId   int       `json:"schedule_id" gorm:"index"`
	ReminderTime time.Time `json:"reminder_time"`
	Method       string    `json:"method"`
	Type         string    `json:"type"`
	IsSent       bool      `json:"is_sent"`
	UserID       int       `json:"user_id" gorm:"index"`
}
