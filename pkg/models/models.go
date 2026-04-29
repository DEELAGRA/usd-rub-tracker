package models

import "time"

type RateModels struct {
	ID         int
	Rate       float64
	Date       time.Time
	Created_at time.Time
}
