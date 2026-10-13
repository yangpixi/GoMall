package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/stock-service/internal/application/command/stock"
	"github.com/yangpixi/GoMall/shared/response"
)

type StockHandler struct {
	createHandler *stock.CreateSKUStockHandler
}

type reserveRequest struct {
	Quantity int64 `json:"quantity" binding:"required"`
	SkuID    int64 `json:"skuId" binding:"required"`
	OrderID  int64 `json:"orderId" binding:"required"`
}

type createRequest struct {
	SkuID        int64 `json:"skuId" binding:"required"`
	AvailableQTY int64 `json:"availableQty" binding:"required"`
	LockedQTY    int64 `json:"lockedQty" binding:"required"`
}

func NewStockHandler(ch *stock.CreateSKUStockHandler) (*StockHandler, error) {
	if ch == nil {
		return nil, errors.New("invalid http stock handler")
	}

	return &StockHandler{createHandler: ch}, nil
}

func (h *StockHandler) CreateHandler(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.createHandler.Handle(c.Request.Context(), &stock.CreateSKUStockCommand{
		SkuID:        req.SkuID,
		AvailableQTY: req.AvailableQTY,
		LockedQTY:    req.LockedQTY,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, "sku stock creation succeeded")
}
