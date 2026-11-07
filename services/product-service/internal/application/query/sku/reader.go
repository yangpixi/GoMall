package sku

import "context"

type Reader interface {
	GetPrice(ctx context.Context, id int64) (*PriceVO, error)
}
