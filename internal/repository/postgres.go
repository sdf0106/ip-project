package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	usersTable      = "users"
	agentsTable     = "agents"
	ownersTable     = "owners"
	clientsTable    = "clients"
	housesTable     = "houses"
	clientCartTable = "client_cart"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const maxAttempts = 3

func NewPostgresDB(cfg Config) *pgxpool.Pool {
	//TODO: get context.Background from app
	ctx := context.Background()

	var dbPool *pgxpool.Pool
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	dbPool, err = pgxpool.Connect(ctx, dsn)

	//utils.DoWithTries(func() error {
	//	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//	defer cancel()
	//
	//	dbPool, err = pgxpool.Connect(ctx, dsn)
	//
	//	if err != nil {
	//		logrus.Fatalf("error on connection to db: %s", err.Error())
	//		return err
	//	}
	//
	//	return nil
	//}, maxAttempts, 5*time.Second)

	err = dbPool.Ping(ctx)
	if err != nil {
		return nil
	}

	return dbPool
}
