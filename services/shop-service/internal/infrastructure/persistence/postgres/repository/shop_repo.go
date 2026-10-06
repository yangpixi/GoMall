package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/yangpixi/GoMall/services/shop-service/internal/domain/shop"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/mapper"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type ShopRepo struct {
	db *gorm.DB
}

func NewShopRepo(db *gorm.DB) shop.Repository {
	return &ShopRepo{db: db}
}

func (r *ShopRepo) FindByID(ctx context.Context, id int64) (*shop.Shop, error) {
	s, err := gorm.G[*model.Shop](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shop.ErrShopNotFound
		}
	}

	sh, err := mapper.ToShop(s)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer shop, id: %d, error: %w", id, err)
	}

	return sh, nil
}

func (r *ShopRepo) FindByOwnerID(ctx context.Context, oID int64) (*shop.Shop, error) {
	s, err := gorm.G[*model.Shop](r.db).Where("owner_id = ?", oID).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shop.ErrShopNotFound
		}
	}

	sh, err := mapper.ToShop(s)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer shop, owner_id: %d, error: %w", oID, err)
	}

	return sh, nil
}

func (r *ShopRepo) Save(ctx context.Context, s *shop.Shop) error {
	po, err := mapper.ToShopPO(s)
	if err != nil {
		return fmt.Errorf("failed to save shop: %w", err)
	}

	err = gorm.G[model.Shop](r.db).Create(ctx, po)
	if err != nil {
		return fmt.Errorf("failed to save shop: %w", err)
	}

	return nil

}
