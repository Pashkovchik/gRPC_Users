package service

import (
	"gRPC_Users/internal/dto"
	"gRPC_Users/internal/entity"
	"gRPC_Users/internal/repository"
)

type User interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	DeleteUser(userId dto.UserId) error
	GetUser() ([]entity.User, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository, repoRedis *repository.Redis) *Service {
	return &Service{
		User: NewUserService(repos.User, repoRedis.RedisUsers),
	}
}
