package models

import "time"

type RateModels struct {
	ID        int       `json:"ID"`
	Rate      float64   `json:"rate"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"createdAt"`
}

type RateModelsDTO struct {
	Rate float64   `json:"rate"`
	Date time.Time `json:"date"`
}
