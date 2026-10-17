package grpc

import (
	"context"

	stockv1 "github.com/yangpixi/GoMall/api/gen/go/stock/v1"
	"github.com/yangpixi/GoMall/services/stock-service/internal/application/command/stock"
)

type StockServer struct {
	stockv1.UnimplementedStockServiceServer

	reserveHandler *stock.ReserveHandler
	releaseHandler *stock.ReleaseHandler
}

func NewStockServer(reserveHandler *stock.ReserveHandler, releaseHandler *stock.ReleaseHandler) *StockServer {
	return &StockServer{reserveHandler: reserveHandler, releaseHandler: releaseHandler}
}

func (s *StockServer) Reserve(ctx context.Context, in *stockv1.ReserveRequest) (*stockv1.ReserveResponse, error) {
	err := s.reserveHandler.Handle(ctx, &stock.ReserveCommand{
		Quantity: in.Quantity,
		SkuID:    in.SkuID,
		OrderID:  in.OrderID,
	})
	if err != nil {
		return nil, err
	}

	return &stockv1.ReserveResponse{}, nil
}

func (s *StockServer) Release(ctx context.Context, in *stockv1.ReleaseRequest) (*stockv1.ReleaseResponse, error) {
	err := s.releaseHandler.Handle(ctx, &stock.ReleaseCommand{ID: in.Id})
	if err != nil {
		return nil, err
	}

	return &stockv1.ReleaseResponse{}, nil
}
