package profile

import "context"

type Repository interface {
	FindByUserID(ctx context.Context, userID int64) (*Profile, error)
	Save(ctx context.Context, profile *Profile) error
}
