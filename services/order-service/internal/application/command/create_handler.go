package command

import (
	"context"
	"fmt"

	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreateHandler struct {
	repo      order.Repository
	generator order.IDGenerator
}

func NewCreateHandler(r order.Repository, g order.IDGenerator) *CreateHandler {
	return &CreateHandler{repo: r, generator: g}
}

func (h *CreateHandler) Handle(ctx context.Context, cmd *CreateOrderCommand) error {
	userID, ok := id.UserIDFromCtx(ctx)
	if !ok {
		return fmt.Errorf("failed to retrieve userID from context")
	}

	address, err := order.NewAddress(cmd.Address.Address, cmd.Address.Recipient, cmd.Address.Phone)
	if err != nil {
		return err
	}

	items := make([]*order.Item, 0, len(cmd.Items))
	for _, item := range cmd.Items {
		itemID := h.generator.NextID()

		// TODO get price from product service
		unitPrice := int64(0)
		totalPrice := unitPrice * item.Quantity
		i, err := order.NewItem(itemID, item.SkuID, item.Quantity, unitPrice, totalPrice)
		if err != nil {
			return err
		}

		items = append(items, i)
	}

	orderID := h.generator.NextID()

	// TODO count various amount from items included
	o, err := order.NewOrder(orderID, userID, int64(0), int64(0), int64(0), items, address)

	err = h.repo.Save(ctx, o)
	if err != nil {
		return err
	}

	return nil
}
