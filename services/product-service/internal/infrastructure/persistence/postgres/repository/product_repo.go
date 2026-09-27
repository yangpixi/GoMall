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

func NewProductRepo(db *gorm.DB) product.Repository {
	return &ProductRepo{db: db}
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

func (r *ProductRepo) Save(ctx context.Context, p *product.Product) error {
	po, err := mapper.ToProductPO(p)
	if err != nil {
		return fmt.Errorf("failed to save product: %w", err)
	}

	err = gorm.G[model.Product](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save product: %w", err)
	}

	return nil
}
