package reservation

type Status string

const (
	StatusReserved  Status = "reserved"
	StatusConfirmed Status = "confirmed"
	StatusReleased  Status = "released"
)

type StockReservation struct {
	id       int64
	orderID  int64
	skuID    int64
	quantity int64
	status   Status
}
