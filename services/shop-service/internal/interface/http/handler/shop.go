package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/shop-service/internal/application/command"
	"github.com/yangpixi/GoMall/shared/response"
)

type ShopHandler struct {
	createHandler *command.CreateShopHandler
}

type createShopRequest struct {
	OwnerID     int64
	Name        string
	Description string
}

func NewShopHandler(ch *command.CreateShopHandler) (*ShopHandler, error) {
	if ch == nil {
		return nil, errors.New("invalid shop handler")
	}

	return &ShopHandler{createHandler: ch}, nil
}

func (h *ShopHandler) CreateHandler(c *gin.Context) {
	var req createShopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.createHandler.Handle(c.Request.Context(), &command.CreateShopCommand{
		OwnerID:     req.OwnerID,
		Name:        req.Name,
		Description: req.Description,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, "shop creation successfully")

}
