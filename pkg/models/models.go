package models

import "time"

type RateModels struct {
	ID        int
	Rate      float64
	Date      time.Time
	CreatedAt time.Time
}
