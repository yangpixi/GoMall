package model

import "time"

type Role struct {
	ID        int64 `gorm:"primarykey"`
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Role) TableName() string {
	return "gomall_role"
}
