package dto

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"time"
)

// LoginRequest defines model for LoginRequest.
type LoginInitRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginResponseData defines model for LoginResponseData.
type LoginInitResponseData struct {
	AccessToken           *string               `json:"access_token,omitempty"`
	AccessTokenExpiresAt  *time.Time            `json:"access_token_expires_at,omitempty"`
	RefreshToken          *string               `json:"refresh_token,omitempty"`
	RefreshTokenExpiresAt *time.Time            `json:"refresh_token_expires_at,omitempty"`
	SessionId             *openapi_types.UUID   `json:"session_id,omitempty"`
	User                  *RegisterResponseData `json:"user,omitempty"`
}
