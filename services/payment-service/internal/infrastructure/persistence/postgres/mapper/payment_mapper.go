package mapper

import (
	"github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/persistence/postgres/model"
)

func ToPayment(m *model.Payment) (*payment.Payment, error) {
	return payment.Restore(&payment.State{
		UserID:        m.UserID,
		OrderID:       m.OrderID,
		Amount:        m.Amount,
		Currency:      m.Currency,
		Channel:       m.Channel,
		Status:        m.Status,
		TransactionID: m.TransactionID,
		ExpiresAt:     m.ExpiresAt,
		PaidAt:        m.PaidAt,
		CancelledAt:   m.CancelledAt,
	})
}

func ToPaymentPO(p *payment.Payment) *model.Payment {
	s := p.Snapshot()
	return &model.Payment{
		ID:            s.ID,
		UserID:        s.UserID,
		OrderID:       s.OrderID,
		Amount:        s.Amount,
		Currency:      s.Currency,
		Channel:       s.Channel,
		Status:        s.Status,
		TransactionID: s.TransactionID,
		ExpiresAt:     s.ExpiresAt,
		PaidAt:        s.PaidAt,
		CancelledAt:   s.CancelledAt,
	}
}
