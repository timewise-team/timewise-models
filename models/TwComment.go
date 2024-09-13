package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwComment struct {
	gorm.Model
	ScheduleId int       `json:"schedule_id"`
	UserId     int       `json:"user_id" gorm:"index"`
	Commenter  string    `json:"commenter"`
	Content    string    `json:"content"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
