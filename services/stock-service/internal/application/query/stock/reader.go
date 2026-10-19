package stock

import "context"

type Reader interface {
	GetStock(ctx context.Context, skuID int64) (*VO, error)
}
