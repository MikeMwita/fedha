package entity

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	UserId            string `json:"user_id" `
	UserName          string `json:"username" `
	Email             string `json:"email" `
	PhoneNumber       string `json:"phone_number"`
	Role              string `json:"role" validate:"required"`
	Hash              string `json:"_"`
	PasswordHash      string `json:"password_hash" `
	Password          string `json:"password" `
	CreatedAt         *time.Time
	FullName          string `json:"full_name" `
	PasswordChangedAt *time.Time
	UserID            uuid.UUID
}

func (u *User) SanitizePassword() {
	u.Password = ""
}

func (u *User) HashPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) PrepareCreate() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	if u.Role != "" {
		u.Role = strings.ToLower(strings.TrimSpace(u.Role))
	}

	return nil
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
