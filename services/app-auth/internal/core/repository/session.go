package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"time"
)

const (
	basePrefix = "api-session:"
)

type sessionRepo struct {
	redisClient *redis.Client
	basePrefix  string
	cfg         *config.Config
}

func (s *sessionRepo) CreateSession(ctx context.Context, sess *entity.Session, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.CreateSession")
	defer span.Finish()

	sess.SessionID = uuid.New().String()
	sessionKey := s.createKey(sess.SessionID)

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		return "", errors.WithMessage(err, "sessionRepo.CreateSession.json.Marshal")
	}
	if err = s.redisClient.Set(ctx, sessionKey, sessBytes, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", errors.Wrap(err, "sessionRepo.CreateSession.redisClient.Set")
	}
	return sessionKey, nil
}

// Get session by id

func (s *sessionRepo) GetSessionByID(ctx context.Context, sessionID string) (*entity.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.GetSessionByID")
	defer span.Finish()

	sessBytes, err := s.redisClient.Get(ctx, sessionID).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "sessionRep.GetSessionByID.redisClient.Get")
	}

	sess := &entity.Session{}
	if err = json.Unmarshal(sessBytes, &sess); err != nil {
		return nil, errors.Wrap(err, "sessionRepo.GetSessionByID.json.Unmarshal")
	}
	return sess, nil
}

// Delete session by id

func (s *sessionRepo) DeleteByID(ctx context.Context, sessionID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.DeleteByID")
	defer span.Finish()

	if err := s.redisClient.Del(ctx, sessionID).Err(); err != nil {
		return errors.Wrap(err, "sessionRepo.DeleteByID")
	}
	return nil
}

func (s *sessionRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.basePrefix, sessionID)
}

func (s *sessionRepo) Invalidate() dto.DefaultRes[string] {
	ctx := context.Background()

	// Get all session keys from Redis.
	keys, err := s.redisClient.Keys(ctx, "*").Result()
	if err != nil {
		return dto.DefaultRes[string]{
			Error: "Error getting session keys from Redis.",
			Code:  500,
		}
	}

	// Delete all session keys from Redis.
	if err := s.redisClient.Del(ctx, keys...).Err(); err != nil {
		return dto.DefaultRes[string]{
			Error: "Error deleting session keys from Redis.",
			Code:  500,
		}
	}
	return dto.DefaultRes[string]{
		Message: "All sessions invalidated successfully.",
	}
}

func NewSessionRepository(redisClient *redis.Client, cfg *config.Config) adapters.SessionRepo {
	return &sessionRepo{redisClient: redisClient, basePrefix: basePrefix, cfg: cfg}
}
