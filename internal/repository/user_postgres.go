package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
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
	table, err := getTableName(role)
	if err != nil {
		return err
	}
	query1 := fmt.Sprintf("INSERT INTO %s (user_id) VALUES($1)", table)
	query2 := fmt.Sprintf("UPDATE %s SET user_type=$1 WHERE id=$2", usersTable)

	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), query1, userId)

	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	_, err = tx.Exec(context.Background(), query2, role, userId)

	if err != nil {
		tx.Rollback(context.Background())
		return err
	}

	tx.Commit(context.Background())

	return err
}

func (r *UserPostgres) UpdateUserInfo(userId int, user domain.User) (domain.User, error) {
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
	query := fmt.Sprintf("UPDATE users SET %s WHERE id=$%d", setQuery, argId)
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

func (r *UserPostgres) ChangeRole(userId int, prevRole string, role string) error {
	prevRoleTable, err := getTableName(prevRole)
	if err != nil {
		return err
	}

	roleTable, err := getTableName(role)
	if err != nil {
		return err
	}

	query1 := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", prevRoleTable)
	query2 := fmt.Sprintf("INSERT INTO %s (user_id) VALUES($1)", roleTable)
	query3 := fmt.Sprintf("UPDATE %s AS ut SET ut.user_type=$1 WHERE ut.id=$2", usersTable)

	tx, err := r.db.Begin(context.Background())

	_, err = tx.Exec(context.Background(), query1, userId)

	if err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback(context.Background())
		return err
	}

	_, err = tx.Exec(context.Background(), query2, userId)

	if err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback(context.Background())
		return err
	}

	_, err = tx.Exec(context.Background(), query3, role, userId)

	if err != nil {
		// Rollback the transaction in case of an error
		tx.Rollback(context.Background())
		return err
	}

	err = tx.Commit(context.Background())

	return err
}

func getTableName(role string) (string, error) {
	switch role {
	case "client":
		return clientsTable, nil
	case "owner":
		return ownersTable, nil
	case "agent":
		return agentsTable, nil
	default:
		return role, errors.New("invalid type of user")
	}
}
