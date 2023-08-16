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
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *AuthRedisRepository) DeleteSession(ctx context.Context, id string) error {
	return r.client.Del(ctx, id).Err()
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
