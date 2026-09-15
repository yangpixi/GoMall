package model

import "time"

type Profile struct {
	UserID    int `gorm:"primaryKey"`
	Nickname  string
	Phone     string
	Email     string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
