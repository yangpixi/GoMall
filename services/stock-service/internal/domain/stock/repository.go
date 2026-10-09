package stock

import "context"

type Repository interface {
	FindBySkuID(ctx context.Context, skuID int64) (*Stock, error)
	Save(ctx context.Context, stock *Stock) error
}
