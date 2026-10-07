package grpc

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	shopv1 "github.com/yangpixi/GoMall/api/gen/go/shop/v1"
	"github.com/yangpixi/GoMall/services/shop-service/internal/application/query/shop"
	"github.com/yangpixi/GoMall/shared/http/id"
	"google.golang.org/grpc/metadata"
)

type ShopServer struct {
	shopv1.UnimplementedShopServiceServer

	handler *shop.GetShopIDHandler
}

func NewShopServer(h *shop.GetShopIDHandler) *ShopServer {
	return &ShopServer{handler: h}
}

func (s *ShopServer) GetShopID(ctx context.Context, req *shopv1.GetShopIDRequest) (*shopv1.GetShopIDResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to retrieve metadata from rpc context")
	}

	if len(md.Get("userID")) == 0 {
		return nil, errors.New("missing required userID")
	}

	userID, err := strconv.ParseInt(md.Get("userID")[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to transform userID from string")
	}

	ctx = id.WithUserID(context.Background(), userID)
	vo, err := s.handler.Handle(ctx)
	if err != nil {
		return nil, err
	}

	shopID, err := strconv.ParseInt(vo.ShopID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to transform shopID from string")
	}

	return &shopv1.GetShopIDResponse{ShopId: shopID}, nil
}
