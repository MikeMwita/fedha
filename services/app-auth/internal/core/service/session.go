package service

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type SessionService struct {
	repo adapters.AuthRepository
}

func (d SessionService) CreateSession(ctx context.Context, session *entity.Session, expire int) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d SessionService) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (d SessionService) DeleteByID(ctx context.Context, sessionID string) error {
	//TODO implement me
	panic("implement me")
}

func (d SessionService) Invalidate() dto.DefaultRes[string] {
	return dto.DefaultRes[string]{
		Message: "The session has been invalidated successfully",
		Error:   "",
		Code:    200,
		Data:    "",
	}
}

type SessionRecord struct {
	UserId string
	Id     string
}

func (d SessionService) ClearSession(background context.Context, id string) error {
	// Get the session record for the user ID
	//sessionRecord := &SessionRecord{UserId: id}
	////err := d.repo.DeleteSession(background, sessionRecord)
	//if err != nil {
	//	return err
	//}

	return nil
}
func NewSessionService(repo adapters.AuthRepository) adapters.SessionService {
	return &SessionService{repo: repo}
}
