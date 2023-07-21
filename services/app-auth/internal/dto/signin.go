package dto

import "github.com/MikeMwita/fedha.git/services/app-auth/internal/core /entity"

type Signin struct {
	User     string ` json:"user" binding:"required"`
	Password string ` json:"password" binding:"required"`
}

type SignInRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type SignInResult struct {
	StatusCode int32  `json:"-"`
	Message    string `json:"message,omitempty"`
}

type LoginOTPBody struct {
	OtpCode string `json:"otp_code" binding:"required"`
}

type ResendOTPReq struct {
	TrackingUID string `json:"tracking_uid"`
}

type ResendOTPRes struct {
	StatusCode   int32  `json:"-"`
	Message      string `json:"message,omitempty"`
	TrackingUuid string `json:"tracking_uuid,omitempty"`
}

type LoginRes struct {
	StatusCode   int32        `json:"-"`
	Message      string       `json:"message,omitempty"`
	Token        string       `json:"bearer_token,omitempty"`
	User         *entity.User `json:"user"`
	RefreshToken string       `json:"refresh_token"`
}

type UserLoginRes struct {
	entity.User
	Token string
}
