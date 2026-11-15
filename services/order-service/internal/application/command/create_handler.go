package command

import (
	"context"
	"fmt"

	"github.com/yangpixi/GoMall/services/order-service/internal/application/port"
	"github.com/yangpixi/GoMall/services/order-service/internal/domain/order"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type CreateHandler struct {
	repo        order.Repository
	generator   order.IDGenerator
	priceReader port.PriceReader
}

func NewCreateHandler(r order.Repository, g order.IDGenerator, pr port.PriceReader) *CreateHandler {
	return &CreateHandler{repo: r, generator: g, priceReader: pr}
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
	orderTotalPrice := int64(0)
	for _, item := range cmd.Items {
		itemID := h.generator.NextID()

		info, err := h.priceReader.GetPrice(ctx, item.SkuID)
		if err != nil {
			return err
		}

		totalPrice := info.Price * item.Quantity
		orderTotalPrice += totalPrice
		i, err := order.NewItem(itemID, item.SkuID, item.Quantity, info.Price, totalPrice)
		if err != nil {
			return err
		}

		items = append(items, i)
	}

	orderID := h.generator.NextID()

	discount := int64(0) // temporarily not count for discount
	o, err := order.NewOrder(orderID, userID, orderTotalPrice, discount, orderTotalPrice-discount, items, address)

	err = h.repo.Save(ctx, o)
	if err != nil {
		return err
	}

	return nil
}
