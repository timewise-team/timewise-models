package user_email_dtos

type SearchUserEmailResponse struct {
	Id             int    `json:"id"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
}

type UserEmailStatusResponse struct {
	Id             int    `json:"id"`
	Email          string `json:"email"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	ProfilePicture string `json:"profile_picture"`
	Status         string `json:"status"`
	IsVerified     bool   `json:"is_verified"`
}
