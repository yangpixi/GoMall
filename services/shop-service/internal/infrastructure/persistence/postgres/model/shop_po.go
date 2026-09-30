package model

import "time"

type Shop struct {
	ID          int64 `gorm:"primaryKey"`
	OwnerID     int64
	Name        string
	Description string
	IsBanned    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
