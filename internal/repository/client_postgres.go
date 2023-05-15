package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClientPostgres struct {
	db *pgxpool.Pool
}

func NewClientPostgres(db *pgxpool.Pool) *ClientPostgres {
	return &ClientPostgres{
		db: db,
	}
}
