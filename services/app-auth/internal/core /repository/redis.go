package repository

import "github.com/go-redis/redis/v8"

type AuthRedisRepository struct {
	client *redis.Client
}

func NewAuthRedisRepository(host string) *AuthRedisRepository {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &AuthRedisRepository{
		client: client,
	}
}
