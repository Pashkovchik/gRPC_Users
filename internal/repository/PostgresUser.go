package repository

import (
	"fmt"
	"gRPC_Users/internal/dto"
	"gRPC_Users/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PostgresUser struct {
	db *sqlx.DB
}

func NewPostgresUser(db *sqlx.DB) *PostgresUser {
	return &PostgresUser{db: db}
}

func (p *PostgresUser) CreateUser(userProfile dto.UserProfile) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO #{userTable} (name, email) VALUES (&1, &2) RETURNING id")
	row := p.db.QueryRow(query, userProfile.Name, userProfile.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PostgresUser) GetUsers() ([]entity.User, error) {
	var users []entity.User

	query := fmt.Sprintf("SELECT u.id, u.name, u.emeil FROM #{usersTable} u")
	err := p.db.Select(&users, query)

	return users, err
}

func (p *PostgresUser) DeleteUser(userId dto.UserId) error {
	query := fmt.Sprintf("DELETE FROM #{usersTable} u WHERE u.id = &1")
	_, err := p.db.Exec(query, userId.Id)

	return err
}
