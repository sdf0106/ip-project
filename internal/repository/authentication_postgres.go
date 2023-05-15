package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type AuthenticationPostgres struct {
	db *pgxpool.Pool
}

func NewAuthenticationPostgres(db *pgxpool.Pool) *AuthenticationPostgres {
	return &AuthenticationPostgres{
		db: db,
	}
}

func (r *AuthenticationPostgres) CreateUser(user domain.User) (int, error) {
	query := fmt.Sprintf("INSERT  INTO %s (name, email, password_hash, registered_at) values ($1, $2, $3, $4) RETURNING ID", usersTable)
	row := r.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password, user.RegisteredAt)

	if err := row.Scan(&user.Id); err != nil {
		return user.Id, err
	}

	return user.Id, nil
}

func (r *AuthenticationPostgres) GetUser(email, password string) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	row := r.db.QueryRow(context.Background(), query, email, password)

	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.RegisteredAt)

	return user, err
}
