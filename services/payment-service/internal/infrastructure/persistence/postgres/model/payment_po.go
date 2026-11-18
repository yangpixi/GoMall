package model

import (
	"time"

	"github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
)

type Payment struct {
	ID            int64 `gorm:"primaryKey"`
	UserID        int64
	OrderID       int64
	Amount        int64
	Currency      string
	Channel       string
	Status        payment.Status
	TransactionID *string

	ExpiresAt   time.Time
	PaidAt      *time.Time // nullable
	CancelledAt *time.Time // nullable

	CreatedAt time.Time
	UpdatedAt time.Time
}
