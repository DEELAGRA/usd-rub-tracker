package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"usd-rub-tracker/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Handler {
	return &Handler{pool: pool}
}

func (h *Handler) Routes(r chi.Router) {
	r.Get("/api/v1/rates/current", h.GetLastRateHandler)
	r.Get("/api/v1/rates/history", h.GetAllRateHandler)
	r.Get("/", h.HelloHandler)
}

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "hello world"
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Printf("Не смог отправить Hello: %v", err)
	}

}
func (h *Handler) GetLastRateHandler(w http.ResponseWriter, r *http.Request) {
	var rate models.RateModelsDTO
	row := h.pool.QueryRow(r.Context(), "SELECT rate, date FROM rates ORDER BY id DESC LIMIT 1")
	err := row.Scan(&rate.Rate, &rate.Date)
	if err != nil {
		log.Printf("Ошибка: Unable to execute query: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)

}
func (h *Handler) GetAllRateHandler(w http.ResponseWriter, r *http.Request) {

	daysSTR := r.URL.Query().Get("days")
	if daysSTR == "" {
		msg := "Прислали пустой квери параметр days"
		w.Write([]byte(msg))
		return
	}
	days, err := strconv.Atoi(daysSTR)
	if err != nil {
		msg := "invalid query parameter 'days', only accept 'int' " + err.Error()
		w.Write([]byte(msg))
		return
	}

	if days < 1 || days > 365 {
		msg := "invalid query parameter 'days', from 1 to 365"
		w.Write([]byte(msg))
		return
	}

	sqlQuery := `
	SELECT rate, date
	FROM rates
	WHERE date >= NOW() - INTERVAL '1 day' * $1
	ORDER BY date DESC
	`
	row, err := h.pool.Query(r.Context(), sqlQuery, days)
	if err != nil {
		log.Printf("Ошибка: Unable to execute query: %v", err)
	}
	defer row.Close()

	var rate []models.RateModelsDTO
	for row.Next() {
		var r models.RateModelsDTO
		if err := row.Scan(&r.Rate, &r.Date); err != nil {
			log.Printf("Scan failed: %v", err)
			msg := "Scan failed: " + err.Error()
			w.Write([]byte(msg))
			return
		}
		rate = append(rate, r)

	}

	if err := row.Err(); err != nil {
		log.Printf("Iteration error: %v", err)
		msg := "iteration error:" + err.Error()
		w.Write([]byte(msg))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}
