package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type AuthRedisRepository struct {
	client *redis.Client
}

func (r *AuthRedisRepository) SetAccessToken(ctx context.Context, key string, value string, expiration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (r *AuthRedisRepository) DeleteSession(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthRedisRepository(host string) *AuthRedisRepository {

	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &AuthRedisRepository{
		client: client,
	}
}

func (r *AuthRedisRepository) SaveAccessToken(ctx context.Context, userID, accessToken string) error {
	return r.client.Set(ctx, userID, accessToken, 0).Err()
}

func (r *AuthRedisRepository) GetAccessToken(ctx context.Context, userID string) (string, error) {
	return r.client.Get(ctx, userID).Result()
}

func (r *AuthRedisRepository) DeleteAccessToken(ctx context.Context, userID string) error {
	return r.client.Del(ctx, userID).Err()
}
