package reservation

import (
	"context"
)

type Repository interface {
	FindByID(ctx context.Context, id int64) (*StockReservation, error)
	FindByOrderID(ctx context.Context, orderID int64) (*StockReservation, error)
	FindBySkuID(ctx context.Context, skuID int64) (*StockReservation, error)
	Save(ctx context.Context, sr *StockReservation) error
}
