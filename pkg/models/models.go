package models

import "time"

type RateModels struct {
	ID        int       `json:"ID"`
	Rate      float64   `json:"Rate"`
	Date      time.Time `json:"Date"`
	CreatedAt time.Time `json:"CreatedAt"`
}

type RateModelsDTO struct {
	Rate float64   `json:"Rate"`
	Date time.Time `json:"Date"`
}
