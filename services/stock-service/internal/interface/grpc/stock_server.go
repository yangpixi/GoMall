package grpc

import (
	"context"

	stockv1 "github.com/yangpixi/GoMall/api/gen/go/stock/v1"
	"github.com/yangpixi/GoMall/services/stock-service/internal/application/command/stock"
)

type StockServer struct {
	stockv1.UnimplementedStockServiceServer

	handler *stock.ReserveHandler
}

func NewStockServer(h *stock.ReserveHandler) *StockServer {
	return &StockServer{handler: h}
}

func (s *StockServer) Reserve(ctx context.Context, in *stockv1.ReserveRequest) (*stockv1.ReserveResponse, error) {
	err := s.handler.Handle(ctx, &stock.ReserveCommand{
		Quantity: in.Quantity,
		SkuID:    in.SkuID,
		OrderID:  in.OrderID,
	})
	if err != nil {
		return nil, err
	}

	return &stockv1.ReserveResponse{}, nil
}
