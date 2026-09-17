package profile

import "context"

type Repository interface {
	FindByUserID(ctx context.Context, userID int64) (*Profile, error)
}
