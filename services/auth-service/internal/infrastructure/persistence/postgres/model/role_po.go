package model

import "time"

type Role struct {
	ID        uint `gorm:"primarykey"`
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Role) TableName() string {
	return "gomall_role"
}
