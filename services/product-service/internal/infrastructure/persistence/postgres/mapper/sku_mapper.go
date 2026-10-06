package mapper

import (
	"github.com/yangpixi/GoMall/services/product-service/internal/domain/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/model"
)

func ToSKU(s *model.SKU) (*sku.SKU, error) {
	return sku.Restore(&sku.State{
		ID:            s.ID,
		ProductID:     s.ProductID,
		Price:         s.Price,
		Specification: s.Specification,
	})
}

func ToSkuPO(s *sku.SKU) (*model.SKU, error) {
	po, err := s.Snapshot()
	if err != nil {
		return nil, err
	}

	return &model.SKU{
		ID:            po.ID,
		ProductID:     po.ProductID,
		Price:         po.Price,
		Specification: po.Specification,
	}, nil
}
