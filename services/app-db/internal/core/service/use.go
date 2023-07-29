package service

import "github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"

type UserService struct {
	userRepo ports.UserRepo
}
