package payment

import "time"

type CallbackCommand struct {
	TradeNo      string
	PaymentID    int64
	Status       string
	CallbackTime time.Time
}
