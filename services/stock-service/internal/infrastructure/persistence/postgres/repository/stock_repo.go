package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/stock"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type StockRepo struct {
	db *gorm.DB
}

func NewStockRepo(db *gorm.DB) *StockRepo {
	return &StockRepo{db: db}
}

func (r *StockRepo) FindBySkuID(ctx context.Context, skuID int64) (*stock.Stock, error) {
	ss, err := gorm.G[*model.SKUStock](r.db).Where("sku_id = ?", skuID).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, stock.ErrSkuStockNotFound
		}
		return nil, fmt.Errorf("failed to select sku stock, skuID: %d, error: %w", skuID, err)
	}

	return mapper.ToStock(ss)
}

func (r *StockRepo) Save(ctx context.Context, stock *stock.Stock) error {
	po, err := mapper.ToStockPO(stock)
	if err != nil {
		return fmt.Errorf("failed to transfer to po: %w", err)
	}

	err = gorm.G[model.SKUStock](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save stock: %w", err)
	}

	return nil
}
