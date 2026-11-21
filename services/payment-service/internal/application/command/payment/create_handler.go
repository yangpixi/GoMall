package payment

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yangpixi/GoMall/services/payment-service/internal/application/port"
	"github.com/yangpixi/GoMall/services/payment-service/internal/application/port/gateway"
	paymentdomain "github.com/yangpixi/GoMall/services/payment-service/internal/domain/payment"
	paymentinfra "github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/payment"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreatePaymentHandler struct {
	repo        paymentdomain.Repository
	generator   paymentdomain.IDGenerator
	registry    *paymentinfra.Registry
	orderReader port.OrderReader
}

func NewCreatePaymentHandler(
	r paymentdomain.Repository,
	generator paymentdomain.IDGenerator,
	registry *paymentinfra.Registry,
	reader port.OrderReader,
) *CreatePaymentHandler {
	return &CreatePaymentHandler{repo: r, generator: generator, orderReader: reader, registry: registry}
}

func (h *CreatePaymentHandler) Handle(ctx context.Context, cmd *CreatePaymentCommand) error {
	o, err := h.orderReader.GetOrderInfo(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	if o.Status != "PendingPayment" {
		return paymentdomain.ErrCannotPay
	}

	pID := h.generator.NextID()

	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return fmt.Errorf("failed to retrieve userID from context")
	}

	p, err := paymentdomain.New(pID, userID, cmd.OrderID, cmd.Amount, cmd.Currency, cmd.Channel)
	if err != nil {
		return err
	}

	err = h.repo.Save(ctx, p)

	if err != nil {
		return err
	}

	// invoke payment gateway
	g, err := h.registry.Get(cmd.Channel)
	if err != nil {
		return err
	}

	_, err = g.Prepay(ctx, &gateway.PrepayRequest{
		OutTradeNo:  strconv.FormatInt(pID, 10),
		Description: fmt.Sprintf("GoMall: %d", pID),
		Amount:      o.FinalAmount,
	})
	if err != nil {
		return err
	}

	return nil
}
