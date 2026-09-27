package model

import "time"

type SKU struct {
	ID            int64 `gorm:"primaryKey"`
	ProductID     int64
	Price         int
	Specification string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
