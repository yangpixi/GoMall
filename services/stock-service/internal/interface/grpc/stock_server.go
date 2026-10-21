package grpc

import (
	"context"

	stockv1 "github.com/yangpixi/GoMall/api/gen/go/stock/v1"
	stockcommand "github.com/yangpixi/GoMall/services/stock-service/internal/application/command/stock"
	stockquery "github.com/yangpixi/GoMall/services/stock-service/internal/application/query/stock"
)

type StockServer struct {
	stockv1.UnimplementedStockServiceServer

	reserveHandler  *stockcommand.ReserveHandler
	releaseHandler  *stockcommand.ReleaseHandler
	confirmHandler  *stockcommand.ConfirmHandler
	getStockHandler *stockquery.GetStockHandler
}

func NewStockServer(
	reserve *stockcommand.ReserveHandler,
	release *stockcommand.ReleaseHandler,
	confirm *stockcommand.ConfirmHandler,
	getStock *stockquery.GetStockHandler,
) *StockServer {
	return &StockServer{
		reserveHandler:  reserve,
		releaseHandler:  release,
		confirmHandler:  confirm,
		getStockHandler: getStock,
	}
}

func (s *StockServer) Reserve(ctx context.Context, in *stockv1.ReserveRequest) (*stockv1.ReserveResponse, error) {
	err := s.reserveHandler.Handle(ctx, &stockcommand.ReserveCommand{
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
	err := s.releaseHandler.Handle(ctx, &stockcommand.ReleaseCommand{ID: in.Id})
	if err != nil {
		return nil, err
	}

	return &stockv1.ReleaseResponse{}, nil
}

func (s *StockServer) Confirm(ctx context.Context, in *stockv1.ConfirmRequest) (*stockv1.ConfirmResponse, error) {
	err := s.confirmHandler.Handle(ctx, &stockcommand.ConfirmCommand{ID: in.Id})
	if err != nil {
		return nil, err
	}

	return &stockv1.ConfirmResponse{}, nil
}

func (s *StockServer) GetStock(ctx context.Context, in *stockv1.GetStockRequest) (*stockv1.GetStockResponse, error) {
	vo, err := s.getStockHandler.Handle(ctx, &stockquery.GetStockQuery{SkuID: in.SkuId})
	if err != nil {
		return nil, err
	}

	return &stockv1.GetStockResponse{
		SkuID:        vo.SkuID,
		AvailableQty: vo.AvailableQTY,
	}, nil
}
