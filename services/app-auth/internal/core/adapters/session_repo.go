package adapters

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

//go:generate mockgen -source session_repo.go -destination repository/mock/redis_repository_mock.go -package mock

type SessionRepo interface {
	Invalidate() dto.DefaultRes[string]
	CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}
