package adapters

import "github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"

type AuthService interface {
	Register(request dto.RegisterReq) (*dto.RegisterRes, error)
	Login(request dto.LoginInitRequest) (*dto.LoginInitResponseData, error)
}

type SessionService interface {
	Invalidate() dto.DefaultRes[string]
}
