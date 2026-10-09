package mapper

import (
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/stock"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/model"
)

func ToStock(m *model.SKUStock) (*stock.Stock, error) {
	return stock.Restore(&stock.State{
		SkuID:        m.SkuID,
		AvailableQTY: m.AvailableQTY,
		LockedQTY:    m.LockedQTY,
	})
}

func ToStockPO(s *stock.Stock) (*model.SKUStock, error) {
	snap, err := s.Snapshot()
	if err != nil {
		return nil, err
	}

	return &model.SKUStock{
		SkuID:        snap.SkuID,
		AvailableQTY: snap.AvailableQTY,
		LockedQTY:    snap.LockedQTY,
	}, nil
}
