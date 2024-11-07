package models

import "time"

type TwDocument struct {
	ID            int             `gorm:"primary_key" `
	FileName      string          `gorm:"type:varchar(255)" json:"file_name"`
	FilePath      string          `gorm:"type:varchar(255)" json:"file_path"`
	FileSize      int             `gorm:"type:int" json:"file_size"`
	FileType      string          `gorm:"type:varchar(255)" json:"file_type"`
	IsDeleted     bool            `gorm:"type:tinyint(1)" json:"is_deleted"`
	ScheduleId    int             `gorm:"index" json:"schedule_id"`
	UploadedBy    int             `gorm:"index" json:"workspace_user_id"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     *time.Time      `gorm:"default:null" json:"deleted_at"`
	UploadedAt    time.Time       `json:"uploaded_at"`
	DownloadUrl   string          `gorm:"type:varchar(255)"json:"download_url"`
	WorkspaceUser TwWorkspaceUser `gorm:"foreignkey:UploadedBy;association_foreignkey:ID"`
	Schedule      TwSchedule      `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
}
