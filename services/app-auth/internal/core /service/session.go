package service

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core /adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type DefaultSessionService struct {
	repo adapters.AuthRepo
}

func (d DefaultSessionService) Invalidate() dto.DefaultRes[interface{}] {
	//TODO implement me
	panic("implement me")
}

func NewDefaultSessionService(repo adapters.AuthRepo) adapters.SessionService {
	return &DefaultSessionService{repo: repo}
}
