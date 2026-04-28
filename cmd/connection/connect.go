package connection

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatConnecton(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, "postgres://postgres:admin@localhost:5433/rates")
}
