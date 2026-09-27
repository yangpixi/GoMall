package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/product-service/internal/domain/product"
	"github.com/yangpixi/GoMall/product-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/product-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func (r *ProductRepo) FindByID(ctx context.Context, id int64) (*product.Product, error) {
	p, err := gorm.G[*model.Product](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, product.ErrProductNotFound
		}

		return nil, fmt.Errorf("failed to select product, id: %w, error: %w", id, err)
	}

	var skuIDs []int64

	err = r.db.WithContext(ctx).Model(model.SKU{}).Where("product_id = ?", id).Pluck("id", skuIDs).Error

	if err != nil {
		return nil, fmt.Errorf("failed to select product: id: %w, error: %w", id, err)
	}

	return mapper.ToProduct(p, skuIDs)
}
