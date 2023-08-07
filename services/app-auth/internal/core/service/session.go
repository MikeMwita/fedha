package service

import (
	"context"
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

type SessionRecord struct {
	UserId string
	Id     string
}

func (d DefaultSessionService) ClearSession(background context.Context, id string) error {
	// Get the session record for the user ID
	sessionRecord := &SessionRecord{UserId: id}
	err := d.repo.DeleteSession(background, sessionRecord)
	if err != nil {
		return err
	}

	return nil
}
func NewDefaultSessionService(repo adapters.AuthRepo) adapters.SessionService {
	return &DefaultSessionService{repo: repo}
}
