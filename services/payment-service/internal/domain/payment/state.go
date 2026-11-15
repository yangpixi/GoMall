package payment

import (
	"strings"
	"time"
)

type State struct {
	UserID        int64
	OrderID       int64
	Amount        int64
	Currency      string
	Channel       string
	Status        Status
	TransactionID string // uuid from payment platform
	ExpiresAt     time.Time
	PaidAt        *time.Time // nullable
	CancelledAt   *time.Time // nullable
}

func Restore(s *State) (*Payment, error) {
	if s.UserID == 0 || s.OrderID == 0 || s.Amount < 0 {
		return nil, ErrInvalidPayment
	}

	if strings.TrimSpace(s.Currency) == "" || strings.TrimSpace(s.Channel) == "" || strings.TrimSpace(s.TransactionID) == "" {
		return nil, ErrInvalidPayment
	}

	return &Payment{
		userID:        s.UserID,
		orderID:       s.OrderID,
		amount:        s.Amount,
		currency:      s.Currency,
		channel:       s.Channel,
		status:        s.Status,
		transactionID: s.TransactionID,
		expiresAt:     s.ExpiresAt,
		paidAt:        s.PaidAt,
		cancelledAt:   s.CancelledAt,
	}, nil
}
