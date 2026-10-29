package model

import (
	"time"

	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
)

type GomallOrder struct {
	ID             int64 `gorm:"primaryKey"`
	UserID         int64
	Status         order.Status
	TotalAmount    int64
	DiscountAmount int64
	FinalAmount    int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
