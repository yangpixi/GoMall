package rpc

import (
	"context"
	"errors"
	"strconv"

	shopv1 "github.com/yangpixi/GoMall/api/gen/go/shop/v1"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/port"
	"github.com/yangpixi/GoMall/shared/http/id"
	"google.golang.org/grpc/metadata"
)

type ShopClient struct {
	client shopv1.ShopServiceClient
}

func NewShopClient(c shopv1.ShopServiceClient) *ShopClient {
	return &ShopClient{client: c}
}

func (c *ShopClient) GetShopID(ctx context.Context) (*port.ShopInfo, error) {
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return nil, errors.New("failed to retrieve userID from context")
	}

	ogCtx := metadata.AppendToOutgoingContext(ctx, "userID", strconv.FormatInt(userID, 10))
	rep, err := c.client.GetShopID(ogCtx, &shopv1.GetShopIDRequest{UserId: userID})
	if err != nil {
		return nil, err
	}

	return &port.ShopInfo{ShopID: rep.ShopId}, nil
}
