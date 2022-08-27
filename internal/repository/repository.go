package repository

import (
	"gRPC_Users/internal/dto"
	"gRPC_Users/internal/entity"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(userProfile dto.UserProfile) (int, error)
	GetUsers() ([]entity.User, error)
	DeleteUser(userId dto.UserId) error
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewPostgresUser(db),
	}
}
