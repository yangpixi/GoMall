package stock

import (
	"context"
	"errors"
)

type GetStockQuery struct {
	SkuID int64
}

type GetStockHandler struct {
	reader Reader
}

func NewGetStockHandler(r Reader) (*GetStockHandler, error) {
	if r == nil {
		return nil, errors.New("invalid get stock handler")
	}

	return &GetStockHandler{reader: r}, nil
}

func (h *GetStockHandler) Handle(ctx context.Context, query *GetStockQuery) (*VO, error) {
	return h.reader.GetStock(ctx, query.SkuID)
}
