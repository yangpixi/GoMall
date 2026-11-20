package port

import "context"

type OrderInfo struct {
	ID             int64
	UserID         int64
	Status         string
	TotalAmount    int64
	DiscountAmount int64
	FinalAmount    int64
}

type OrderReader interface {
	GetOrderInfo(ctx context.Context, id int64) (*OrderInfo, error)
}
