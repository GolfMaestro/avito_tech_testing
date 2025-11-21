package tests

import (
	"avito_tech_testing/repository"
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

// run tests: go test ./... -v

func TestConnection(t *testing.T) {

	ctx := context.Background()

	config, err := pgxpool.ParseConfig("postgres://postgres:2004@db:5433/avito_tech_test")
	if err != nil {
		t.Fatal(err)
	}
	config.MaxConns = 5
	config.MinConns = 1

	repository.Pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
}
