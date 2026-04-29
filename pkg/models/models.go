package models

import "time"

type RateModels struct {
	ID         int
	Rates      float64
	Date       time.Time
	Created_at time.Time
}
