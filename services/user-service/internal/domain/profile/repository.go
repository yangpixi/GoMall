package profile

import "context"

type Repository interface {
	FindByUserID(ctx context.Context, userID uint) (*Profile, error)
}
