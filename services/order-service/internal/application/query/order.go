package query

import "context"

type OrderQuery struct {
	ID int64
}

type OrderQueryHandler struct {
	reader Reader
}

func NewOrderQueryHandler(r Reader) *OrderQueryHandler {
	return &OrderQueryHandler{reader: r}
}

func (h *OrderQueryHandler) Handle(ctx context.Context, query *OrderQuery) (*VO, error) {
	vo, err := h.reader.GetOrder(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return vo, nil
}
