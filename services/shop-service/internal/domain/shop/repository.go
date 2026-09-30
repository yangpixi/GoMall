package shop

import "context"

type Repository interface {
	FindByID(ctx context.Context, id int64) (*Shop, error)
	FindByOwnerID(ctx context.Context, oID int64) (*Shop, error)
	Save(ctx context.Context, s *Shop) error
}
