package dto

import "github.com/MikeMwita/fedha.git/services/app-auth/internal/core /entity"

// SignUpJSONBody defines parameters for SignUp.
type SignUpJSONBody struct {
	ConfirmPassword string `json:"confirm_password"`

	// Email valid email
	Email string `json:"email"`

	// Name full user name
	Name string `json:"name"`

	// Password strong password
	Password string `json:"password"`
}

type UserRegistrationRes struct {
	entity.User
	StatusCode int `json:"status_code"`
	Message    string
}
type AccountVerification struct {
	entity.User
	Token        string
	RefreshToken string
}
