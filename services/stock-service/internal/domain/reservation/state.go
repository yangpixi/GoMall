package reservation

import "time"

type State struct {
	ID        int64
	OrderID   int64
	SkuID     int64
	Quantity  int64
	Status    Status
	ExpiredAt time.Time
}

func Restore(s *State) (*StockReservation, error) {
	if s.ID == 0 || s.OrderID == 0 || s.SkuID == 0 || s.Quantity <= 0 || time.Until(s.ExpiredAt) <= 0 {
		return nil, ErrInvalidStockReservation
	}

	return &StockReservation{
		id:        s.ID,
		orderID:   s.OrderID,
		skuID:     s.SkuID,
		quantity:  s.Quantity,
		status:    s.Status,
		expiredAt: s.ExpiredAt,
	}, nil
}
