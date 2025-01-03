package models

import (
	"time"
)

type TwUserEmail struct {
	ID         int        `gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"default:null"`
	UserId     int        `json:"user_id" gorm:"index"`
	Email      string     `json:"email"`
	Status     *string    `json:"status"`
	IsLinkedTo *int       `json:"is_linked_to"`
	ExpiresAt  *time.Time `json:"expires_at"`
	User       TwUser     `gorm:"foreignkey:UserId;association_foreignkey:ID"`
}
