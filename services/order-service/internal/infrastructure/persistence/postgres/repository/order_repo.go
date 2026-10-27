package repository

import (
	"context"

	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) order.Repository {
	return &OrderRepo{db: db}
}

func (o *OrderRepo) FindByID(ctx context.Context, id int64) (*order.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepo) Save(ctx context.Context, order *order.Order) error {
	//TODO implement me
	panic("implement me")
}
