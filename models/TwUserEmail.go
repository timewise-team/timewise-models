package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwUserEmail struct {
	gorm.Model
	UserId    int       `json:"user_id" gorm:"index"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}