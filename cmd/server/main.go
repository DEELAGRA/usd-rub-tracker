package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	database "usd-rub-tracker/internal/db"

	"github.com/go-chi/chi/v5"
)

func main() {

	ctx := context.Background()
	pool, err := database.CreatConnection(ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("БД недоступна\n %v", err)
	}
	fmt.Println("db connect!")
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

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Ошибка запуска сервера %v\n", err)
	}

}
