package model

import "time"

type UserAddress struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Address   string
	Phone     string
	Recipient string
	CreatedAt time.Time
	UpdatedAt time.Time
}
