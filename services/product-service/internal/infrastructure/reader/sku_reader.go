package reader

import (
	"context"
	"errors"
	"fmt"

	skuquery "github.com/yangpixi/GoMall/services/product-service/internal/application/query/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/domain/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type SKUReader struct {
	db *gorm.DB
}

func NewSKUReader(db *gorm.DB) *SKUReader {
	return &SKUReader{db: db}
}

func (r *SKUReader) GetPrice(ctx context.Context, id int64) (*skuquery.PriceVO, error) {
	s, err := gorm.G[*model.SKU](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sku.ErrSKUNotFound
		}
		return nil, fmt.Errorf("failed to select sku, id: %d, error: %w", id, err)
	}

	return &skuquery.PriceVO{Price: s.Price}, nil
}
