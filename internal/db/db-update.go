package database

import (
	"context"
	"fmt"
	models "usd-rub-tracker/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SaveRate(ctx context.Context, pool *pgxpool.Pool, rate models.RateModels) error {
	sqlQuery := `
 	INSERT INTO rates(rate, date, created_at)
 	VALUES($1,$2,$3)
	ON CONFLICT (date)
	DO UPDATE SET rate = EXCLUDED.rate, created_at = EXCLUDED.created_at
 	`
	_, err := pool.Exec(ctx, sqlQuery, rate.Rate, rate.Date, rate.CreatedAt)
	if err != nil {
		return fmt.Errorf("SaveRate: не удалось записать данные: %w", err)
	}
	return nil
}
