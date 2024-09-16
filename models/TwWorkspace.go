package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TwWorkspace struct {
	gorm.Model
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ExtraData   string    `json:"extra_data"`
	Description string    `json:"description"`
	CreatedBy   int       `json:"created_by"`
	Key         string    `json:"key"`
	Type        string    `json:"type"`
	IsDeleted   bool      `json:"is_deleted"`
	DeletedAt   time.Time `json:"deleted_at"`
}