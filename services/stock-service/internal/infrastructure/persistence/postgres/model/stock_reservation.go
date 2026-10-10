package model

import (
	"time"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
)

type StockReservation struct {
	ID        int64 `gorm:"primaryKey"`
	OrderID   int64
	SkuID     int64
	Quantity  int64
	Status    reservation.Status
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
