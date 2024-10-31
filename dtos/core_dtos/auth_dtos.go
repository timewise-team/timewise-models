package core_dtos

type GoogleAuthRequest struct {
	Credentials string `json:"credentials"`
}

type GoogleAuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	IsNewUser   bool   `json:"is_new_user"`
	IdToken     string `json:"id_token"`
}
