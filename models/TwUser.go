package models

import (
	"time"
)

type TwUser struct {
	ID                   int       `gorm:"primary_key"`
	FirstName            string    `gorm:"type:varchar(100)"`
	LastName             string    `gorm:"type:varchar(100)"`
	Email                string    `gorm:"type:varchar(255)"`
	PasswordHash         string    `gorm:"type:varchar(255)"`
	ProfilePicture       string    `gorm:"type:varchar(255)"`
	Timezone             string    `gorm:"type:varchar(100)"`
	Locale               string    `gorm:"type:varchar(10)"`
	IsVerified           bool      `gorm:"type:tinyint(1)"`
	IsActive             bool      `gorm:"type:tinyint(1)"`
	CreatedAt            time.Time `gorm:"type:datetime"`
	UpdatedAt            time.Time `gorm:"type:datetime"`
	LastLoginAt          time.Time `gorm:"type:datetime"`
	NotificationSettings string    `gorm:"type:text"`
	CalendarSettings     string    `gorm:"type:text"`
	Role                 string    `gorm:"type:varchar(50)"`
}
