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

func (r *UserPostgres) ChangeRole(userId int, prevRole string, role string) error {
	prevRoleTable, err := getTableName(prevRole)
	if err != nil {
		return err
	}

	roleTable, err := getTableName(role)
	if err != nil {
		return err
	}

	query1 := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", prevRoleTable)
	query2 := fmt.Sprintf("INSERT INTO %s (user_id) VALUES($1)", roleTable)

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
