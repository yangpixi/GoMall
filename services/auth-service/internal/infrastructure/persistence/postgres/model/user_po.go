package model

import (
	"time"

	"github.com/yangpixi/GoMall/auth-service/internal/domain/account"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Status    account.Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "gomall_user"
}
