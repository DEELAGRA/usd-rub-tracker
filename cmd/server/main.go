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

	"github.com/go-chi/chi/middleware"
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

	cbr.FetchUSDRAteSave(ctx, pool)
	ticker := time.NewTicker(6 * time.Hour)
	go func() {
		for range ticker.C {
			cbr.FetchUSDRAteSave(ctx, pool)
		}
	}()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	ratesHandler := handlers.New(pool)
	ratesHandler.Routes(r)
	fs := http.FileServer(http.Dir("./frontend"))
	r.Mount("/", http.StripPrefix("/", fs))

	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Printf("Ошибка запуска сервера %v\n", err)
	}

}
