package reader

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/order-service/internal/application/query"
	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type OrderReader struct {
	db *gorm.DB
}

func NewOrderReader(db *gorm.DB) *OrderReader {
	return &OrderReader{db: db}
}

func (r *OrderReader) GetOrder(ctx context.Context, id int64) (*query.VO, error) {
	po, err := gorm.G[*model.GomallOrder](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrOrderNotFound
		}
		return nil, fmt.Errorf("failed to select order, id: %d, error: %w", id, err)
	}

	return &query.VO{
		ID:             po.ID,
		UserID:         po.UserID,
		Status:         po.Status,
		TotalAmount:    po.TotalAmount,
		DiscountAmount: po.DiscountAmount,
		FinalAmount:    po.FinalAmount,
	}, nil
}
