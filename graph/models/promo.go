package models

import "time"

type Promo struct {
	ID       int64 `gorm:"primaryKey"`
	Discount float64
	EndAt    time.Time
}
