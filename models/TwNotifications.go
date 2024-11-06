package models

import "time"

type TwNotifications struct {
	ID              int         `gorm:"primary_key"`
	UserEmailId     int         `json:"user_email_id" gorm:"index"`
	Type            string      `json:"type"`
	Message         string      `json:"message"`
	IsRead          bool        `json:"is_read"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	DeletedAt       *time.Time  `json:"deleted_at" gorm:"default:null"`
	RelatedItemId   int         `json:"related_item_id"`
	RelatedItemType string      `json:"related_item_type"`
	ExtraData       string      `json:"extra_data"`
	UserEmail       TwUserEmail `gorm:"foreignkey:UserEmailId ;association_foreignkey:ID"`
	IsSent          bool        `json:"is_sent" gorm:"default:false"`
	NotifiedAt      *time.Time  `json:"notified_at" gorm:"default:CURRENT_TIMESTAMP(3)"`
}
