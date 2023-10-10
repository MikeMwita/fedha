package adapters

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
)

//go:generate mockgen -source session_service.go -destination mock/session_service_mock.go -package mock

type SessionService interface {
	CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
