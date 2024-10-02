package models

import (
	"time"
)

type TwUserEmail struct {
	ID        int       `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UserId    int       `json:"user_id" gorm:"index"`
	Email     string    `json:"email"`
	IsLinked  bool      `json:"is_linked"`
}
