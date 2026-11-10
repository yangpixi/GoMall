package grpc

import (
	"context"

	orderv1 "github.com/yangpixi/GoMall/api/gen/go/order/v1"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/command"
)

type OrderServer struct {
	orderv1.UnimplementedOrderServiceServer

	createHandler *command.CreateHandler
}

func NewOrderServer(ch *command.CreateHandler) *OrderServer {
	return &OrderServer{
		createHandler: ch,
	}
}

func (s *OrderServer) Create(ctx context.Context, in *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {
	itemCMDs := make([]*command.ItemCommand, 0, len(in.Items))
	for _, item := range in.Items {
		itemCMDs = append(itemCMDs, &command.ItemCommand{
			SkuID:    item.SkuId,
			Quantity: item.Quantity,
		})
	}

	err := s.createHandler.Handle(ctx, &command.CreateOrderCommand{
		Items: itemCMDs,
		Address: &command.AddressCommand{
			Address:   in.Address.Address,
			Recipient: in.Address.Recipient,
			Phone:     in.Address.Phone,
		},
	})

	if err != nil {
		return nil, err
	}

	return &orderv1.CreateOrderResponse{Msg: "success"}, nil
}
