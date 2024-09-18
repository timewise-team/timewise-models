package user_login_dtos

type UserLoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"PasswordHash"`
}
