package payment

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Payment, error)
	Save(ctx context.Context, payment *Payment) error
}
