package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDBConnetion() error {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig("postgres://postgres:2004@localhost:5432/avito_tech")
	if err != nil {
		return err
	}
	config.MaxConns = 5
	config.MinConns = 1

	Pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return err
	}

	return nil
}
