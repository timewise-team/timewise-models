package models

import (
	"time"
)

type TwUserEmail struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UserId    int       `json:"user_id" gorm:"index"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
