package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/command"
	"github.com/yangpixi/GoMall/shared/response"
)

type OrderHandler struct {
	createHandler *command.CreateHandler
}

func NewOrderHandler(ch *command.CreateHandler) *OrderHandler {
	return &OrderHandler{createHandler: ch}
}

type item struct {
	SkuID    int64
	Quantity int64
}

type address struct {
	address   string
	recipient string
	phone     string
}

type createOrderRequest struct {
	items   []*item
	address *address
}

func (h *OrderHandler) CreateHandler(c *gin.Context) {
	req := createOrderRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	itemCMDs := make([]*command.ItemCommand, 0, len(req.items))
	for _, item := range req.items {
		itemCMDs = append(itemCMDs, &command.ItemCommand{
			SkuID:    item.SkuID,
			Quantity: item.Quantity,
		})
	}

	err := h.createHandler.Handle(c.Request.Context(), &command.CreateOrderCommand{
		Items: itemCMDs,
		Address: &command.AddressCommand{
			Address:   req.address.address,
			Recipient: req.address.recipient,
			Phone:     req.address.phone,
		},
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, "success")

}
