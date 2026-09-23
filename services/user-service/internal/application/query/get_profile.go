package query

import "context"

type GetProfileQuery struct {
	UserID int64 // only for admin side
}

type ProfileVO struct {
	UserID   string `json:"userId"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type ProfileReader interface {
	GetProfileDetail(ctx context.Context, userID int64) (*ProfileVO, error)
}
