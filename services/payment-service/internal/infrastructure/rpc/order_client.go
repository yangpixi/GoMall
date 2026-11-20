package rpc

import (
	"context"

	orderv1 "github.com/yangpixi/GoMall/api/gen/go/order/v1"
	"github.com/yangpixi/GoMall/services/payment-service/internal/application/port"
)

type OrderClient struct {
	client orderv1.OrderServiceClient
}

func NewOrderClient(c orderv1.OrderServiceClient) *OrderClient {
	return &OrderClient{client: c}
}

func (c *OrderClient) GetOrderInfo(ctx context.Context, id int64) (*port.OrderInfo, error) {
	res, err := c.client.Query(ctx, &orderv1.OrderQueryRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &port.OrderInfo{
		ID:             res.Id,
		UserID:         res.UserId,
		Status:         res.Status,
		TotalAmount:    res.TotalAmount,
		DiscountAmount: res.DiscountAmount,
		FinalAmount:    res.FinalAmount,
	}, nil
}
