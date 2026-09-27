package mapper

import (
	"github.com/yangpixi/GoMall/product-service/internal/domain/sku"
	"github.com/yangpixi/GoMall/product-service/internal/infrastructure/persistence/postgres/model"
)

func ToSKU(s *model.SKU) (*sku.SKU, error) {
	return sku.Restore(&sku.State{
		ID:            s.ID,
		ProductID:     s.ProductID,
		Price:         s.Price,
		Specification: s.Specification,
	})
}
