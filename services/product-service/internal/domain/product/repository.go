package product

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Product, error)
	Save(ctx context.Context, product *Product) error
}
