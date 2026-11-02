package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) order.Repository {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) FindByID(ctx context.Context, id int64) (*order.Order, error) {
	o, err := gorm.G[*model.GomallOrder](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrOrderNotFound
		}
		return nil, fmt.Errorf("failed to select order, id: %d, error: %w", id, err)
	}

	add, err := gorm.G[*model.GomallOrderAddress](r.db).Where("order_id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrAddressNotFound
		}
		return nil, fmt.Errorf("failed to select order address, order_id: %d, error: %w", id, err)
	}

	items, err := gorm.G[*model.GomallOrderItem](r.db).Where("order_id = ?", id).Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to select order items, order_id: %d, error: %w", id, err)
	}

	if len(items) == 0 {
		return nil, order.ErrItemNotFound
	}

	return mapper.ToOrder(o, items, add)
}

func (r *OrderRepo) Save(ctx context.Context, order *order.Order) error {
	orderPO := mapper.ToOrderPO(order)
	err := gorm.G[model.GomallOrder](r.db).Create(ctx, orderPO)
	if err != nil {
		return fmt.Errorf("failed to save order po")
	}

	addressPO := mapper.ToAddressPO(order.Address(), order.ID())
	err = gorm.G[model.GomallOrderAddress](r.db).Create(ctx, addressPO)
	if err != nil {
		return fmt.Errorf("failed to save address po")
	}

	batchItem := make([]*model.GomallOrderItem, 0, len(order.Items()))
	for _, item := range order.Items() {
		itemPO := mapper.ToItemPO(item, order.ID())
		batchItem = append(batchItem, itemPO)
	}

	err = gorm.G[*model.GomallOrderItem](r.db).CreateInBatches(ctx, &batchItem, len(batchItem))
	if err != nil {
		return fmt.Errorf("failed to save item pos")
	}

	return nil
}
