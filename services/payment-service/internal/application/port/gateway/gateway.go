package gateway

import "context"

type PrepayRequest struct {
	OutTradeNo  string
	Description string
	Amount      int64
}

type PrepayResponse struct {
	Result string
}

type QueryRequest struct {
	// TODO
}

type QueryResponse struct {
	// TODO
}

type PaymentGateway interface {
	Prepay(ctx context.Context, req *PrepayRequest) (*PrepayResponse, error)
	Query(ctx context.Context, req *QueryRequest) (*QueryResponse, error)
}
