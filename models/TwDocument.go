package models

type TwDocument struct {
	ID          int         `gorm:"primary_key"`
	FileName    string      `gorm:"type:varchar(255)"`
	FilePath    string      `gorm:"type:varchar(255)"`
	FileSize    int         `gorm:"type:int"`
	FileType    string      `gorm:"type:varchar(255)"`
	IsDeleted   bool        `gorm:"type:tinyint(1)"`
	ScheduleId  int         `gorm:"index"`
	UserId      int         `gorm:"index"`
	CreatedAt   string      `gorm:"type:datetime"`
	UpdatedAt   string      `gorm:"type:datetime"`
	DeletedAt   string      `gorm:"type:datetime" gorm:"default:null"`
	UploadedAt  string      `gorm:"type:datetime"`
	DownloadUrl string      `gorm:"type:varchar(255)"`
	UserEmail   TwUserEmail `gorm:"foreignkey:UserId;association_foreignkey:ID"`
	Schedule    TwSchedule  `gorm:"foreignkey:ScheduleId;association_foreignkey:ID"`
}
