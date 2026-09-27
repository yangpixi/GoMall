package mapper

import (
	"github.com/yangpixi/GoMall/product-service/internal/domain/product"
	"github.com/yangpixi/GoMall/product-service/internal/infrastructure/persistence/postgres/model"
)

func ToProduct(p *model.Product, skuIDs []int64) (*product.Product, error) {
	return product.Restore(&product.State{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Status:      p.Status,
		SkuIDs:      skuIDs,
	})
}

func ToProductPO(p *product.Product) (*model.Product, error) {
	s, err := p.Snapshot()
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          s.ID,
		Name:        s.Name,
		ShopID:      s.ShopID,
		Description: s.Description,
		Status:      s.Status,
	}, nil
}
