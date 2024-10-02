package models

import (
	"time"
)

type TwWorkspace struct {
	ID          int       `gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"default:null"`
	Title       string    `json:"title"`
	ExtraData   string    `json:"extra_data"`
	Description string    `json:"description"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	IsDeleted   bool      `json:"is_deleted"`
}
