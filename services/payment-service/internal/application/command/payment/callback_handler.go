package payment

import (
	"context"
	"errors"

	"github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
)

type CallbackHandler struct {
	repo payment.Repository
}

func NewCallbackHandler(r payment.Repository) *CallbackHandler {
	return &CallbackHandler{repo: r}
}

func (h *CallbackHandler) Handle(ctx context.Context, cmd *CallbackCommand) error {
	if cmd.Status != "success" {
		return errors.New("pay result unsuccessful")
	}

	p, err := h.repo.FindByID(ctx, cmd.PaymentID)
	if err != nil {
		return err
	}

	p.ToStatusPaid(cmd.CallbackTime, cmd.TradeNo)
	err = h.repo.Update(ctx, p)
	if err != nil {
		return err
	}

	return nil
}
