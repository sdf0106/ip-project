package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type UserPostgres struct {
	db *pgxpool.Pool
}

func NewUserPostgres(db *pgxpool.Pool) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) ChooseRole(userId int, role string) error {
	var table string

	switch role {
	case "client":
		table = clientsTable
	case "owner":
		table = ownersTable
	case "agent":
		table = agentsTable
	default:
		return errors.New("invalid type of user")
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id) VALUES($1)", table)
	_, err := r.db.Exec(context.Background(), query, userId)

	return err
}

func (r *UserPostgres) UpdateUserInfo(userId int, user domain.User) (domain.User, error) {
	return user, nil
}

func (r *UserPostgres) ChangeRole(userId int, role string) error {
	return nil
}
