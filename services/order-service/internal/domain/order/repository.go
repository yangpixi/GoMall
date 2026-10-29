package order

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Order, error)
	Save(ctx context.Context, order *Order) error
}
