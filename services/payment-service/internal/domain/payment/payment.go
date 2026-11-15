package payment

import (
	"strings"
	"time"
)

type Payment struct {
	userID        int64
	orderID       int64
	amount        int64
	currency      string
	channel       string
	status        Status
	transactionID string // uuid from payment platform
	expiresAt     time.Time
	paidAt        *time.Time // nullable
	cancelledAt   *time.Time // nullable
}

func New(userID, orderID, amount int64, currency, channel, txID string) (*Payment, error) {
	if userID == 0 || orderID == 0 || amount < 0 {
		return nil, ErrInvalidPayment
	}

	if strings.TrimSpace(currency) == "" || strings.TrimSpace(channel) == "" || strings.TrimSpace(txID) == "" {
		return nil, ErrInvalidPayment
	}

	return &Payment{
		userID:        userID,
		orderID:       orderID,
		amount:        amount,
		currency:      currency,
		channel:       channel,
		status:        StatusPending,
		transactionID: txID,
		expiresAt:     time.Now().Add(30 * time.Minute),
		paidAt:        nil,
		cancelledAt:   nil,
	}, nil
}

func (p *Payment) ToStatusPaid(t time.Time) {
	p.status = StatusPaid
	p.paidAt = &t
}

func (p *Payment) ToStatusCancelled(t time.Time) {
	p.status = StatusCancelled
	p.cancelledAt = &t
}
