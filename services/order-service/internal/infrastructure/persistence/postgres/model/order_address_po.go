package model

import "time"

type GomallOrderAddress struct {
	ID        int64 `gorm:"primaryKey"`
	OrderID   int64
	Address   string
	Recipient string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
