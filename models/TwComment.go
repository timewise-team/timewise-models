package models

import (
	"time"
)

type TwComment struct {
	ID         int       `gorm:"primary_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	ScheduleId int       `json:"schedule_id"`
	UserId     int       `json:"user_id" gorm:"index"`
	Commenter  string    `json:"commenter"`
	Content    string    `json:"content"`
	IsDeleted  bool      `json:"is_deleted"`
}
