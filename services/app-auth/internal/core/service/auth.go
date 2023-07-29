package service

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"

	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type AuthService struct {
	repo   adapters.AuthRepo
	config config.Jwt
}

func (a AuthService) Register(request dto.RegisterReq) (*dto.RegisterRes, error) {
	registerResponse, err := a.Register(request)
	if err != nil {
		return nil, err
	}
	return registerResponse, err
	//panic("implement me")
}

func (a AuthService) Login(request dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	//TODO implement me
	loginResponse, err := a.Login(request)
	if err != nil {
		return nil, err
	}
	return loginResponse, err
}

func NewAuthService(repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo}
}
