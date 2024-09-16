package models

import (
	"time"
)

type TwRecurrenceException struct {
	ID            uint      `gorm:"primary_key"`
	ScheduleId    int       `json:"schedule_id" gorm:"index"`
	ExceptionDate time.Time `json:"exception_date"`
	NewStartTime  time.Time `json:"new_start_time"`
	NewEndTime    time.Time `json:"new_end_time"`
	IsCancelled   bool      `json:"is_cancelled"`
	ExtraData     string    `json:"extra_data"`
}
