package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/stock-service/internal/application/command"
)

type StockHandler struct {
	reserveHandler *command.ReserveHandler
}

type reserveRequest struct {
	Quantity int64 `json:"quantity" binding:"required"`
	SkuID    int64 `json:"skuId" binding:"required"`
	OrderID  int64 `json:"orderId" binding:"required"`
}

func NewStockHandler(rh *command.ReserveHandler) (*StockHandler, error) {
	if rh == nil {
		return nil, errors.New("invalid http stock handler")
	}

	return &StockHandler{reserveHandler: rh}, nil
}

func (h *StockHandler) ReserveHandler(c *gin.Context) {
	var req reserveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.reserveHandler.Handle(c.Request.Context(), &command.ReserveCommand{
		Quantity: req.Quantity,
		SkuID:    req.SkuID,
		OrderID:  req.OrderID,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

}
