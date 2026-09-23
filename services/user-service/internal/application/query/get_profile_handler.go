package query

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/shared/http/id"
)

type GetProfileHandler struct {
	reader ProfileReader
}

func NewGetProfileHandler(r ProfileReader) (*GetProfileHandler, error) {
	if r == nil {
		return nil, errors.New("invalid profile repository")
	}

	return &GetProfileHandler{reader: r}, nil
}

func (h *GetProfileHandler) Handle(ctx context.Context) (*ProfileVO, error) {
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return nil, errors.New("failed to retrieve userID from context")
	}

	return h.reader.GetProfileDetail(ctx, userID)
}
