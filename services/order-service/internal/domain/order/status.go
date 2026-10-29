package order

type Status string

const (
	StatusPaid           Status = "paid"
	StatusCancelled      Status = "cancelled"
	StatusPendingPayment Status = "pending_payment"
)
