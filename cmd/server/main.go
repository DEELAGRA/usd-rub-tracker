package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	cbr "usd-rub-tracker/internal/api/cbr"
	database "usd-rub-tracker/internal/db"

	"github.com/go-chi/chi/v5"
)

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
	/*var rate models.RateModels
	rate.Rate = 16.4
	rate.Date = time.Now()
	rate.Created_at = time.Now()
	if err := database.SaveRate(ctx, pool, rate); err != nil {
		panic(err)
	} */

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world👋"))
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})
	usd, err := cbr.FetchUSDRAte(ctx)
	if err != nil {
		log.Printf("Ошибка при переводе данных: %v", err)
	}
	fmt.Println(usd)

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Ошибка запуска сервера %v\n", err)
	}

}
