package models

import (
	"time"
)

type TwRecurrenceException struct {
	ID            int       `gorm:"primary_key"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	ScheduleId    int       `json:"schedule_id" gorm:"index"`
	ExceptionDate time.Time `json:"exception_date"`
	NewStartTime  time.Time `json:"new_start_time"`
	NewEndTime    time.Time `json:"new_end_time"`
	IsCancelled   bool      `json:"is_cancelled"`
	ExtraData     string    `json:"extra_data"`
}
