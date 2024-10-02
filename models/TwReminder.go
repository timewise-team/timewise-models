package models

import (
	"time"
)

type TwReminder struct {
	ID           int       `gorm:"primary_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	ScheduleId   int       `json:"schedule_id" gorm:"index"`
	ReminderTime time.Time `json:"reminder_time"`
	Method       string    `json:"method"`
	Type         string    `json:"type"`
	IsSent       bool      `json:"is_sent"`
	UserID       int       `json:"user_id" gorm:"index"`
}
