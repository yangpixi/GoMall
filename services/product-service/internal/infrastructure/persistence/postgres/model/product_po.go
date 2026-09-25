package model

import "time"

type Product struct {
	ID          int64 `gorm:"primaryKey"`
	Name        string
	ShopID      int64
	Description string
	Status      int // 1: In stock; 0: Out of stock
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
