package query

import "github.com/yangpixi/GoMall/services/order-service/internal/domain/order"

type VO struct {
	ID             int64
	UserID         int64
	Status         order.Status
	TotalAmount    int64
	DiscountAmount int64
	FinalAmount    int64
}
