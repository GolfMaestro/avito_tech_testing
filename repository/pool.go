package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func InitDBConnetion() error {
	ctx := context.Background()

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	config, err := pgxpool.ParseConfig(conn)
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
