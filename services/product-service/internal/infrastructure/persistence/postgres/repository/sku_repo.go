package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/product-service/internal/domain/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type SKURepo struct {
	db *gorm.DB
}

func NewSKURepo(db *gorm.DB) sku.Repository {
	return &SKURepo{db: db}
}

func (r *SKURepo) FindByID(ctx context.Context, id int64) (*sku.SKU, error) {
	s, err := gorm.G[*model.SKU](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sku.ErrSKUNotFound
		}

		return nil, fmt.Errorf("failed to select sku, id: %d, error: %w", id, err)
	}

	return mapper.ToSKU(s)
}

func (r *SKURepo) FindByProductID(ctx context.Context, pID int64) ([]*sku.SKU, error) {
	pos, err := gorm.G[*model.SKU](r.db).Where("product_id = ?", pID).Find(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sku.ErrSKUNotFound
		}

		return nil, fmt.Errorf("failed to select sku, productID: %d, error: %w", pID, err)
	}

	skus := make([]*sku.SKU, 0, len(pos))

	for _, po := range pos {
		s, err := mapper.ToSKU(po)
		if err != nil {
			return nil, fmt.Errorf("failed to select sku, productID: %d, error: %w", pID, err)
		}
		skus = append(skus, s)
	}

	return skus, nil
}

func (r *SKURepo) Save(ctx context.Context, sku *sku.SKU) error {
	po, err := mapper.ToSkuPO(sku)
	if err != nil {
		return fmt.Errorf("failed to save sku: %w", err)
	}

	err = gorm.G[model.SKU](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save sku: %w", err)
	}

	return nil
}
