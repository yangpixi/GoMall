package profile

import "context"

type Reader interface {
	GetProfileDetail(ctx context.Context, userID int64) (*VO, error)
}
