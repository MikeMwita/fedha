package repository

import "github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"

type userRepo struct {
	db ports.UserStorage
}
