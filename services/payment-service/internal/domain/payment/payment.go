package payment

import (
	"strings"
	"time"
)

type Payment struct {
	id            int64
	userID        int64
	orderID       int64
	amount        int64
	currency      string
	channel       string
	status        Status
	transactionID *string // uuid from payment platform
	expiresAt     time.Time
	paidAt        *time.Time // nullable
	cancelledAt   *time.Time // nullable
}

func New(id, userID, orderID, amount int64, currency, channel string) (*Payment, error) {
	if id == 0 {
		return nil, ErrInvalidPayment
	}

	if userID == 0 || orderID == 0 || amount < 0 {
		return nil, ErrInvalidPayment
	}

	if strings.TrimSpace(currency) == "" || strings.TrimSpace(channel) == "" {
		return nil, ErrInvalidPayment
	}

	return &Payment{
		id:          id,
		userID:      userID,
		orderID:     orderID,
		amount:      amount,
		currency:    currency,
		channel:     channel,
		status:      StatusPending,
		expiresAt:   time.Now().Add(30 * time.Minute),
		paidAt:      nil,
		cancelledAt: nil,
	}, nil
}

func (p *Payment) Snapshot() *State {
	return &State{
		ID:            p.id,
		UserID:        p.userID,
		OrderID:       p.orderID,
		Amount:        p.amount,
		Currency:      p.currency,
		Channel:       p.channel,
		Status:        p.status,
		TransactionID: p.transactionID,
		ExpiresAt:     p.expiresAt,
		PaidAt:        p.paidAt,
		CancelledAt:   p.cancelledAt,
	}
}

func (p *Payment) ToStatusPaid(t time.Time) {
	p.status = StatusPaid
	p.paidAt = &t
}

func (p *Payment) ToStatusCancelled(t time.Time) {
	p.status = StatusCancelled
	p.cancelledAt = &t
}
