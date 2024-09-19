package user_login_dtos

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
