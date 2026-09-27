package sku

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*SKU, error)
	FindByProductID(ctx context.Context, pID int64) ([]*SKU, error)
}
