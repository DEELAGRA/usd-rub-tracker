package cbr

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	database "usd-rub-tracker/internal/db"
	"usd-rub-tracker/pkg/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func FetchUSDRAte(ctx context.Context) (float64, error) {
	const url = "https://www.cbr-xml-daily.ru/daily_json.js"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status: %v", resp.StatusCode)
	}

	var data CBRResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("decode json: %w", err)
	}

	usd, ok := data.Valute["USD"]
	if !ok {
		return 0, fmt.Errorf("Не нашли USD в ответе")
	}

	log.Printf("Получен курс USD: %.4f RUB", usd.Value)
	return usd.Value, nil

}

func FetchUSDRAteSave(ctx context.Context, pool *pgxpool.Pool) {
	usd, err := FetchUSDRAte(ctx)
	if err != nil {
		log.Printf("Ошибка при переводе данных: %v", err)
	}

	rate := models.RateModels{
		Rate:      usd,
		Date:      time.Now(),
		CreatedAt: time.Now(),
	}

	if err := database.SaveRate(ctx, pool, rate); err != nil {
		log.Printf("Ошибка записи в db: %v", err)
	}
}
