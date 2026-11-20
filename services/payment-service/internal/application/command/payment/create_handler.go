package payment

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yangpixi/GoMall/services/payment-service/internal/application/port"
	"github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreatePaymentHandler struct {
	repo        payment.Repository
	generator   payment.IDGenerator
	gateway     port.PaymentGateway
	orderReader port.OrderReader
}

func NewCreatePaymentHandler(
	r payment.Repository,
	generator payment.IDGenerator,
	gateway port.PaymentGateway,
	reader port.OrderReader,
) *CreatePaymentHandler {
	return &CreatePaymentHandler{repo: r, generator: generator, gateway: gateway, orderReader: reader}
}

func (h *CreatePaymentHandler) Handle(ctx context.Context, cmd *CreatePaymentCommand) error {
	o, err := h.orderReader.GetOrderInfo(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	if o.Status != "PendingPayment" {
		return payment.ErrCannotPay
	}

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

	// invoke payment gateway
	_, err = h.gateway.Prepay(ctx, &port.PrepayRequest{
		OutTradeNo:  strconv.FormatInt(pID, 10),
		Description: fmt.Sprintf("GoMall: %d", pID),
		Amount:      o.FinalAmount,
	})
	if err != nil {
		return err
	}

	return nil
}
