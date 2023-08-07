package dto

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"time"
)

type RegisterReq struct {
	Email    openapi_types.Email `json:"email"`
	FullName string              `json:"full_name"`
	Password string              `json:"password"`
	Username string              `json:"username"`
}

type RegisterRes struct {
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	Email             string     `json:"email,omitempty"`
	FullName          string     `json:"full_name,omitempty"`
	PasswordChangedAt *time.Time `json:"password_changed_at,omitempty"`
	Username          string     `json:"username,omitempty"`
}
