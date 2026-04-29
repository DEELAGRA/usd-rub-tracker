package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatConnection(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, "postgres://postgres:admin@localhost:5433/rates")
}
