package model

import "time"

type UserAddress struct {
	UserID    uint `gorm:"primaryKey"`
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
