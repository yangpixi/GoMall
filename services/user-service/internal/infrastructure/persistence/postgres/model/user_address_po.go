package model

import "time"

type UserAddress struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	Address   string
	Phone     string
	Recipient string
	CreatedAt time.Time
	UpdatedAt time.Time
}
