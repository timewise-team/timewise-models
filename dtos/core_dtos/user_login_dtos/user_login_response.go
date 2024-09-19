package user_login_dtos

type UserLoginResponse struct {
	AccessToken string `json:"AccessToken"`
	ExpiresIn   int    `json:"ExpiresIn"`
	TokenType   string `json:"TokenType"`
}
