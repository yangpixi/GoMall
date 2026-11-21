package mock

import (
	"context"
	"log/slog"

	"github.com/yangpixi/GoMall/services/payment-service/internal/application/port/gateway"
	"github.com/yangpixi/GoMall/services/payment-service/internal/config"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/payment"
)

func init() {
	payment.Register("mock", NewMockChannel)
}

type Channel struct {
}

func NewMockChannel(cfg *config.Config) (gateway.PaymentGateway, error) {
	return &Channel{}, nil
}

func (c *Channel) Prepay(ctx context.Context, req *gateway.PrepayRequest) (*gateway.PrepayResponse, error) {
	slog.Info("mock payment channel", "tradeNo", req.OutTradeNo, "Amount", req.Amount)
	return &gateway.PrepayResponse{Result: "success"}, nil
}

func (c *Channel) Query(ctx context.Context, req *gateway.QueryRequest) (*gateway.QueryResponse, error) {
	return &gateway.QueryResponse{}, nil
}
