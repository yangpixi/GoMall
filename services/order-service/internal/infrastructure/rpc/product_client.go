package rpc

import (
	"context"
	"strconv"

	productv1 "github.com/yangpixi/GoMall/api/gen/go/product/v1"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/port"
	"google.golang.org/grpc/metadata"
)

type ProductClient struct {
	client productv1.SKUServiceClient
}

func NewProductClient(c productv1.SKUServiceClient) *ProductClient {
	return &ProductClient{client: c}
}

func (c *ProductClient) GetPrice(ctx context.Context, id int64) (*port.PriceInfo, error) {
	ogCtx := metadata.AppendToOutgoingContext(ctx, "sku_id", strconv.FormatInt(id, 10))
	rep, err := c.client.GetSKUPrice(ogCtx, &productv1.GetSKUPriceRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &port.PriceInfo{Price: rep.Price}, nil
}
