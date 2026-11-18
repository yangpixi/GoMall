package payment

import (
	"context"
	"fmt"

	"github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreatePaymentHandler struct {
	repo      payment.Repository
	generator payment.IDGenerator
}

func NewCreatePaymentHandler(r payment.Repository) *CreatePaymentHandler {
	return &CreatePaymentHandler{repo: r}
}

func (h *CreatePaymentHandler) Handle(ctx context.Context, cmd *CreatePaymentCommand) error {
	pID := h.generator.NextID()

	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return fmt.Errorf("failed to retrieve userID from context")
	}

	p, err := payment.New(pID, userID, cmd.OrderID, cmd.Amount, cmd.Currency, cmd.Channel)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, p)

	if err != nil {
		return err
	}

	return nil
}
