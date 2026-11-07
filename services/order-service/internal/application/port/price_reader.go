package port

import "context"

type PriceInfo struct {
	Price int64
}

type PriceReader interface {
	GetPrice(ctx context.Context, id int64) (*PriceInfo, error)
}
