package grpc

import (
	"context"

	orderv1 "github.com/yangpixi/GoMall/api/gen/go/order/v1"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/query"
)

type OrderServer struct {
	orderv1.UnimplementedOrderServiceServer

	createHandler     *command.CreateHandler
	orderQueryHandler *query.OrderQueryHandler
}

func NewOrderServer(ch *command.CreateHandler, qh *query.OrderQueryHandler) *OrderServer {
	return &OrderServer{
		createHandler:     ch,
		orderQueryHandler: qh,
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

func (s *OrderServer) Query(ctx context.Context, in *orderv1.OrderQueryRequest) (*orderv1.OrderQueryResponse, error) {
	vo, err := s.orderQueryHandler.Handle(ctx, &query.OrderQuery{ID: in.Id})
	if err != nil {
		return nil, err
	}

	return &orderv1.OrderQueryResponse{
		Id:             vo.ID,
		UserId:         vo.UserID,
		Status:         string(vo.Status),
		TotalAmount:    vo.TotalAmount,
		DiscountAmount: vo.DiscountAmount,
		FinalAmount:    vo.FinalAmount,
	}, nil
}
