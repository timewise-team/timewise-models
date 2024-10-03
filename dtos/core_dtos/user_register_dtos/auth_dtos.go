package user_register_dto

import "github.com/timewise-team/timewise-models/models"

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

type GetOrCreateUserRequestDto struct {
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	VerifiedEmail  bool   `json:"verified_email"`
	GivenName      string `json:"given_name"`
	FamilyName     string `json:"family_name"`
	Locale         string `json:"locale"`
	GoogleId       string `json:"google_id"`
}

type GetOrCreateUserResponseDto struct {
	User      models.TwUser `json:"user"`
	IsNewUser bool          `json:"is_new_user"`
}
