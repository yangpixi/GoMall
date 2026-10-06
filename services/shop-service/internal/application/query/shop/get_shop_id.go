package shop

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/shared/http/id"
)

type GetShopIDQuery struct {
	UserID int64 // for admin endpoint
}

type GetShopIDHandler struct {
	reader Reader
}

func NewGetShopIDHandler(r Reader) (*GetShopIDHandler, error) {
	if r == nil {
		return nil, errors.New("invalid get shop id handler")
	}

	return &GetShopIDHandler{reader: r}, nil
}

func (h *GetShopIDHandler) Handle(ctx context.Context) (*VO, error) {
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return nil, errors.New("failed to retrieve userID from context")
	}

	return h.reader.GetShopID(ctx, userID)
}
