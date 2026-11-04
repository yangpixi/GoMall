package grpc

import (
	"context"

	productv1 "github.com/yangpixi/GoMall/api/gen/go/product/v1"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/query/sku"
)

type ProductServer struct {
	productv1.UnimplementedSKUServiceServer

	priceHandler *sku.GetSKUPriceHandler
}

func NewProductServer(ph *sku.GetSKUPriceHandler) *ProductServer {
	return &ProductServer{
		priceHandler: ph,
	}
}

func (s *ProductServer) GetSKUPrice(ctx context.Context, in *productv1.GetSKUPriceRequest) (*productv1.GetSKUPriceResponse, error) {
	vo, err := s.priceHandler.Handle(ctx, &sku.GetSKUPriceQuery{ID: in.Id})
	if err != nil {
		return nil, err
	}

	return &productv1.GetSKUPriceResponse{Price: vo.Price}, nil
}
