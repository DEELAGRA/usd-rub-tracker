package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	cbr "usd-rub-tracker/internal/api/cbr"
	"usd-rub-tracker/internal/api/handlers"
	database "usd-rub-tracker/internal/db"
	"usd-rub-tracker/pkg/models"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
			if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
			next.ServeHTTP(w, r)
	})
}

func main() {

	ctx := context.Background()
	pool, err := database.CreatConnection(ctx)
	if err != nil {
		log.Printf("БД не подключена: %v", err)
	}

	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Printf("БД недоступна\n %v", err)
	}

	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for range ticker.C {
			usd, err := cbr.FetchUSDRAte(ctx)
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
	}()
	

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(corsMiddleware)

	ratesHandler := handlers.New(pool)
	ratesHandler.Routes(r)

	
	
	
	
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Ошибка запуска сервера %v\n", err)
	}

}
