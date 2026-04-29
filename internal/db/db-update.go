package database

import (
	"context"
	models "usd-rub-tracker/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveRate(ctx context.Context, pool *pgxpool.Pool, rates models.RateModels) error {
	sqlQuery := `
 	INSERT INTO rates(rate, date, created_at)
 	VALUES($1,$2,$3)
 	`
	_, err := pool.Exec(ctx, sqlQuery, rates.Rate, rates.Date, rates.Created_at)
	return err
}
