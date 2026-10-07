package port

import (
	"context"
)

type ShopInfo struct {
	ShopID int64
}

type ShopIDReader interface {
	GetShopID(ctx context.Context) (*ShopInfo, error)
}
