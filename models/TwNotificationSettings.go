package models

import "time"

type TwNotificationSettings struct {
	ID                           int       `gorm:"primary_key"`
	UserId                       int       `json:"user_id" gorm:"index"`
	NotificationOnTag            bool      `json:"notification_on_tag"`
	NotificationOnComment        bool      `json:"notification_on_comment"`
	NotificationOnDueDate        bool      `json:"notification_on_due_date"`
	NotificationOnScheduleChange bool      `json:"notification_on_schedule_change"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	DeletedAt                    time.Time `json:"deleted_at" gorm:"default:null"`
	ExtraData                    string    `json:"extra_data"`
	User                         TwUser    `gorm:"foreignkey:UserId;association_foreignkey:ID"`
	NotificationOnEmail          bool      `json:"notification_on_email"`
}
