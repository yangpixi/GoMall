package mapper

import (
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/model"
)

func ToStockReservation(m *model.StockReservation) (*reservation.StockReservation, error) {
	return reservation.Restore(&reservation.State{
		ID:        m.ID,
		OrderID:   m.OrderID,
		SkuID:     m.SkuID,
		Quantity:  m.Quantity,
		Status:    m.Status,
		ExpiredAt: m.ExpiredAt,
	})
}

func ToStockReservationPO(sr *reservation.StockReservation) (*model.StockReservation, error) {
	s, err := sr.Snapshot()
	if err != nil {
		return nil, err
	}

	return &model.StockReservation{
		ID:        s.ID,
		OrderID:   s.OrderID,
		SkuID:     s.SkuID,
		Quantity:  s.Quantity,
		Status:    s.Status,
		ExpiredAt: s.ExpiredAt,
	}, nil
}
