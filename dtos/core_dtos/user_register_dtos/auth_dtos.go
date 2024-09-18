package user_register_dto

type RegisterRequestDto struct {
	UserName        string `json:"username"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterResponseDto struct {
	UserName     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}
