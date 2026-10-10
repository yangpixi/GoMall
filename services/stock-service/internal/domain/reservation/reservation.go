package reservation

import "time"

type Status string

const (
	StatusReserved  Status = "reserved"
	StatusConfirmed Status = "confirmed"
	StatusReleased  Status = "released"
)

type StockReservation struct {
	id        int64
	orderID   int64
	skuID     int64
	quantity  int64
	status    Status
	expiredAt time.Time
}

func New(id, orderID, skuID, quantity int64, expiredAt time.Time) (*StockReservation, error) {
	if id == 0 || orderID == 0 || skuID == 0 || quantity <= 0 || time.Until(expiredAt) <= 0 {
		return nil, ErrInvalidStockReservation
	}

	return &StockReservation{
		id:        id,
		orderID:   orderID,
		skuID:     skuID,
		quantity:  quantity,
		status:    StatusReserved,
		expiredAt: expiredAt,
	}, nil
}

func (r *StockReservation) Snapshot() (*State, error) {
	return &State{
		ID:        r.id,
		OrderID:   r.orderID,
		SkuID:     r.skuID,
		Quantity:  r.quantity,
		Status:    r.status,
		ExpiredAt: r.expiredAt,
	}, nil
}

func (r *StockReservation) ToStatusConfirmed() {
	r.status = StatusConfirmed
}

func (r *StockReservation) ToStatusReleased() {
	r.status = StatusReleased
}
