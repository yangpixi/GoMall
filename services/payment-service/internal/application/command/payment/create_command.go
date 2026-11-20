package payment

type CreatePaymentCommand struct {
	UserID   int64 // admin
	OrderID  int64
	Amount   int64
	Currency string
	Channel  string
}
