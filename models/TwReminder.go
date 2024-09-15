package models

import (
	"time"
)

type TwReminder struct {
	ID           int       `json:"id"`
	ScheduleID   int       `json:"schedule_id"`
	ReminderTime time.Time `json:"reminder_time"`
	Method       string    `json:"method"`
	Type         string    `json:"type"`
	IsSent       bool      `json:"is_sent"`
	UserID       int       `json:"user_id"`
}
