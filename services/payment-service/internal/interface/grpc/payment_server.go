package grpc

import (
	"context"
	"fmt"
	"strconv"

	paymentv1 "github.com/yangpixi/GoMall/api/gen/go/payment/v1"
	commandpayment "github.com/yangpixi/GoMall/services/payment-service/internal/application/command/payment"
	"google.golang.org/grpc/metadata"
)

type PaymentServer struct {
	paymentv1.UnimplementedPaymentServiceServer

	createHandler commandpayment.CreatePaymentHandler
}

func NewPaymentServer(ch commandpayment.CreatePaymentHandler) *PaymentServer {
	return &PaymentServer{createHandler: ch}
}

func (s *PaymentServer) Create(ctx context.Context, in *paymentv1.CreatePaymentRequest) (*paymentv1.CreatePaymentResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve userID from gRPC context")
	}
	userID, err := strconv.ParseInt(md.Get("userID")[0], 10, 64)
	if err != nil {
		return nil, err
	}

	err = s.createHandler.Handle(ctx, &commandpayment.CreatePaymentCommand{
		UserID:   userID,
		OrderID:  in.OrderId,
		Amount:   in.Amount,
		Currency: in.Currency,
		Channel:  in.Channel,
	})
	if err != nil {
		return nil, err
	}

	return &paymentv1.CreatePaymentResponse{Msg: "success"}, nil
}
