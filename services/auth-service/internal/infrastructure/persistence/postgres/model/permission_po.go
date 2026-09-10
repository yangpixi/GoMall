package model

import "time"

type Permission struct {
	ID        uint `gorm:"primaryKey"`
	Code      string
	Resource  string
	Action    string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Permission) TableName() string {
	return "gomall_permission"
}
