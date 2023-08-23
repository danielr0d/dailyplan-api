package models

type SingUp struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	UserRole string `json:"user_role" validate:"required,lte=25"`
}

type SingIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
