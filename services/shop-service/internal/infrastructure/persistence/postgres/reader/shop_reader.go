package reader

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yangpixi/GoMall/services/shop-service/internal/application/query/shop"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/model"
	"gorm.io/gorm"
)

type ShopReader struct {
	db *gorm.DB
}

func NewShopReader(db *gorm.DB) *ShopReader {
	return &ShopReader{db: db}
}

func (r *ShopReader) GetShopID(ctx context.Context, userID int64) (*shop.VO, error) {
	var shopID int64
	err := r.db.WithContext(ctx).Model(model.Shop{}).Where("owner_id = ?", userID).Pluck("id", &shopID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get shop id, id: %d, error: %w", userID, err)
	}

	if shopID == 0 {
		return nil, fmt.Errorf("user's shop not found")
	}

	return &shop.VO{ShopID: strconv.FormatInt(shopID, 10)}, nil
}
