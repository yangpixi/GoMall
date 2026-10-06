package profile

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/shared/http/id"
)

type GetProfileQuery struct {
	UserID int64 // only for admin side
}

type GetProfileHandler struct {
	reader Reader
}

func NewGetProfileHandler(r Reader) (*GetProfileHandler, error) {
	if r == nil {
		return nil, errors.New("invalid profile repository")
	}

	return &GetProfileHandler{reader: r}, nil
}

func (h *GetProfileHandler) Handle(ctx context.Context) (*VO, error) {
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return nil, errors.New("failed to retrieve userID from context")
	}

	return h.reader.GetProfileDetail(ctx, userID)
}
