package reader

import (
	"context"
	"fmt"

	"github.com/yangpixi/GoMall/services/stock-service/internal/application/query/stock"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type StockReader struct {
	db *gorm.DB
}

func NewStockReader(db *gorm.DB) *StockReader {
	return &StockReader{db: db}
}

func (r *StockReader) GetStock(ctx context.Context, skuID int64) (*stock.VO, error) {
	s, err := gorm.G[*model.SKUStock](r.db).Where("sku_id = ?", skuID).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query sku stock, skuID: %d, error: %w", skuID, err)
	}

	return &stock.VO{
		SkuID:        s.SkuID,
		AvailableQTY: s.AvailableQTY,
	}, nil
}
