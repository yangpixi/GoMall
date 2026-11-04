package sku

import "context"

type GetSKUPriceQuery struct {
	ID int64
}

type GetSKUPriceHandler struct {
	reader Reader
}

func NewGetSKUPriceHandler(r Reader) *GetSKUPriceHandler {
	return &GetSKUPriceHandler{reader: r}
}

func (h *GetSKUPriceHandler) Handle(ctx context.Context, query *GetSKUPriceQuery) (*PriceVO, error) {
	vo, err := h.reader.GetPrice(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return vo, nil
}
