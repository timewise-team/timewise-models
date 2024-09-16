package core_dtos

import "time"

type TwRecurrenceExceptionResponseDTO struct {
	ID            uint      `json:"id"`
	ScheduleId    int       `json:"schedule_id"`
	ExceptionDate time.Time `json:"exception_date"`
	NewStartTime  time.Time `json:"new_start_time"`
	NewEndTime    time.Time `json:"new_end_time"`
	IsCancelled   bool      `json:"is_cancelled"`
	ExtraData     string    `json:"extra_data,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TwRecurrenceExceptionCreateDTO struct {
	ScheduleId    int       `json:"schedule_id" binding:"required"`
	ExceptionDate time.Time `json:"exception_date"`
	NewStartTime  time.Time `json:"new_start_time"`
	NewEndTime    time.Time `json:"new_end_time"`
	IsCancelled   bool      `json:"is_cancelled"`
	ExtraData     string    `json:"extra_data,omitempty"`
}

type TwRecurrenceExceptionUpdateDTO struct {
	ExceptionDate *time.Time `json:"exception_date,omitempty"`
	NewStartTime  *time.Time `json:"new_start_time,omitempty"`
	NewEndTime    *time.Time `json:"new_end_time,omitempty"`
	IsCancelled   *bool      `json:"is_cancelled,omitempty"`
	ExtraData     *string    `json:"extra_data,omitempty"`
}

type TwRecurrenceExceptionActionResponseDTO struct {
	ID            uint      `json:"id"`
	ScheduleId    int       `json:"schedule_id"`
	ExceptionDate time.Time `json:"exception_date"`
	NewStartTime  time.Time `json:"new_start_time"`
	NewEndTime    time.Time `json:"new_end_time"`
	IsCancelled   bool      `json:"is_cancelled"`
	ExtraData     string    `json:"extra_data,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
