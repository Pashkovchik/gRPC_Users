package service

import (
	"gRPC_Users/internal/dto"
	"gRPC_Users/internal/entity"
	"gRPC_Users/internal/repository"
)

type UserService struct {
	repo      repository.User
	repoRedis repository.RedisUsers
}

func NewUserService(repo repository.User, repoRedis repository.RedisUsers) *UserService {
	return &UserService{repo: repo, repoRedis: repoRedis}
}

func (u *UserService) CreateUser(userProfile dto.UserProfile) (int, error) {
	return u.repo.CreateUser(userProfile)
}

func (u *UserService) DeleteUser(userId dto.UserId) error {
	return u.repo.DeleteUser(userId)
}

func (u *UserService) GetUsers() ([]entity.User, error) {
	users := u.repoRedis.GetUsers()
	if users != nil {
		return users, nil
	}

	users, err := u.repo.GetUsers()
	if err != nil {
		return nil, err
	}

	u.repoRedis.SaveUsers(users)

	return users, nil
}
