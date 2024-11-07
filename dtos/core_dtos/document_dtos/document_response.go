package document_dtos

type TwDocumentResponse struct {
	ID                  int    `gorm:"primary_key" `
	FileName            string `gorm:"type:varchar(255)" json:"file_name"`
	FilePath            string `gorm:"type:varchar(255)" json:"file_path"`
	FileSize            int    `gorm:"type:int" json:"file_size"`
	FileType            string `gorm:"type:varchar(255)" json:"file_type"`
	IsDeleted           bool   `gorm:"type:tinyint(1)" json:"is_deleted"`
	ScheduleId          int    `gorm:"index" json:"schedule_id"`
	UploadedBy          int    `gorm:"index" json:"workspace_user_id"`
	CreatedAt           string `gorm:"type:datetime" json:"created_at"`
	UpdatedAt           string `gorm:"type:datetime" json:"updated_at"`
	DeletedAt           string `gorm:"type:datetime" gorm:"default:null" json:"deleted_at"`
	UploadedAt          string `gorm:"type:datetime" json:"uploaded_at"`
	DownloadUrl         string `gorm:"type:varchar(255)"json:"download_url"`
	Role                string `json:"role"`
	StatusWorkspaceUser string `json:"status_workspace_user"`
	IsVerified          bool   `json:"is_verified"`
	UserId              int    `json:"user_id" `
	Email               string `json:"email"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	ProfilePicture      string `json:"profile_picture"`
}

type TwUploadDocumentRequest struct {
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
	FileSize    int    `json:"file_size"`
	FileType    string `json:"file_type"`
	ScheduleId  int    `json:"schedule_id"`
	UploadedBy  int    `json:"workspace_user_id"`
	DownloadUrl string `json:"download_url"`
}
