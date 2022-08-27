package repository

import (
	"gRPC_Users/internal/entity"
	"github.com/go-redis/redis/v8"
)

type RedisUsers interface {
	SaveUsers(users []entity.User)
	GetUsers() []entity.User
}

type Redis struct {
	RedisUsers
}

func NewRedisRepository(client *redis.Client) *Redis {
	return &Redis{
		NewRedisUser(client),
	}
}
