package repository

import (
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

	return nil
}

func (r *UserPostgres) UpdateUserInfo(userId int, user domain.User) (domain.User, error) {

	return user, nil
}

func (r *UserPostgres) ChangeRole(userId int, role string) error {

	return nil
}
