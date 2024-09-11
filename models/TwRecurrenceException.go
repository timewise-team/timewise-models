package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwRecurrenceException struct {
	gorm.Model
	ID            int       `json:"id"`
	ScheduleID    int       `json:"schedule_id"`
	ExceptionDate time.Time `json:"exception_date"`
	NewStartTime  time.Time `json:"new_start_time"`
	NewEndTime    time.Time `json:"new_end_time"`
	IsCancelled   bool      `json:"is_cancelled"`
	ExtraData     string    `json:"extra_data"`
}
