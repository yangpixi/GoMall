package model

import "time"

type SKUStock struct {
	SkuID        int64 `gorm:"primaryKey"`
	AvailableQTY int64
	LockedQTY    int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
