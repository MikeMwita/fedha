package service

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type DefaultSessionService struct {
	repo adapters.AuthRepo
}

func (d DefaultSessionService) Invalidate() dto.DefaultRes[string] {
	return dto.DefaultRes[string]{
		Message: "The session has been invalidated successfully",
		Error:   "",
		Code:    200,
		Data:    "",
	}
}

func NewDefaultSessionService(repo adapters.AuthRepo) adapters.SessionService {
	return &DefaultSessionService{repo: repo}
}
