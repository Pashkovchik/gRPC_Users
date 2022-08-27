package repository

import (
	"context"
	"encoding/json"
	"gRPC_Users/internal/entity"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisUser struct {
	client *redis.Client
}

func NewRedisUser(client *redis.Client) *RedisUser {
	return &RedisUser{client: client}
}

func (r *RedisUser) SaveUsers(users []entity.User) {
	serialized, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	r.client.Set(context.Background(), usersKey, serialized, cacheExpires*time.Second)
}

func (r *RedisUser) GetUsers() []entity.User {
	data, err := r.client.Get(context.Background(), usersKey).Result()
	if err != nil {
		return nil
	}

	var users []entity.User

	err = json.Unmarshal([]byte(data), &users)
	if err != nil {
		panic(err)
	}

	return users
}
