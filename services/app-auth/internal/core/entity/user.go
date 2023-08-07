package entity

import "time"

type User struct {
	UserId            string `json:"user_id" `
	UserName          string `json:"username" `
	Email             string `json:"email" `
	PhoneNumber       string `json:"phone_number"`
	Hash              string `json:"_"`
	PasswordHash      string `json:"password_hash" `
	Password          string `json:"password" `
	CreatedAt         *time.Time
	FullName          string `json:"full_name" `
	PasswordChangedAt *time.Time
}
