package service

import (
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/reservation"
	"github.com/yangpixi/GoMall/services/stock-service/internal/domain/stock"
)

func Reserve(s *stock.Stock, rev *reservation.StockReservation) error {
	if s.SkuID() != rev.SkuID() {
		return ErrInvalidSKUID
	}

	err := s.Lock(rev.Quantity())

	if err != nil {
		return err
	}

	return nil
}
