package validation

import (
	"context"
	"errors"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/go-playground/validator/v10"
)

var (
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidFullName    = errors.New("invalid full name")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidUsername    = errors.New("username must only contain letters, numbers, and underscores")
	ErrInvalidEmailFormat = errors.New("invalid email format")
)

func ValidateRegistration(ctx context.Context, data dto.RegisterReq) error {
	validate := validator.New()
	// Validate the struct
	if err := validate.Struct(data); err != nil {
		// Get the first validation error
		err := err.(validator.ValidationErrors)[0]

		// Return a custom error based on the field and tag
		switch err.Field() {
		case "Email":
			return ErrInvalidEmailFormat
		case "FullName":
			return ErrInvalidFullName
		case "Password":
			return ErrInvalidPassword
		case "Username":
			return ErrInvalidUsername
		default:
			return err
		}
	}

	return nil

}

func ValidateLogin(ctx context.Context, data dto.LoginInitRequest) error {
	validate := validator.New()
	// Validate the struct
	if err := validate.Struct(data); err != nil {
		// Get the first validation error
		err := err.(validator.ValidationErrors)[0]

		// Return a custom error based on the field and tag
		switch err.Field() {
		case "Email":
			return ErrInvalidEmail
		case "Password":
			return ErrInvalidPassword
		default:
			return err
		}
	}

	return nil

	//// Validate Email
	//if err := validate.Var(r.Email, "required,email"); err != nil {
	//	return ErrInvalidEmail
	//}
	//
	//// Validate Password (min length)
	//if len(r.Password) < 8 {
	//	return ErrInvalidPassword
	//}
	//
	//return nil
}
