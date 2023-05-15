package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/dto"
	"log"
	"strings"
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

func (r *UserPostgres) UpdateUserInfo(userId int, user dto.UpdateUserInput) (domain.User, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, user.Name)
		argId++
	}
	if user.Email != "" {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, user.Email)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE books SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(context.Background(), query, args...)
	if err != nil {
		log.Println(err)
	}
	var result domain.User
	result.Name = user.Name
	result.Email = user.Email
	return result, nil
}

func (r *UserPostgres) ChangeRole(userId int, role string) error {
	query := fmt.Sprintf("UPDATE %s Set user_type=%s WHERE user_id=$1", usersTable, role)

	_, err := r.db.Exec(context.Background(), query, userId)
	if err != nil {
		log.Println(err)
	}
	return nil
}
