package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/command/address"
	"github.com/yangpixi/GoMall/shared/response"
)

type AddressHandler struct {
	createHandler *address.CreateAddressHandler
	updateHandler *address.UpdateAddressHandler
	deleteHandler *address.DeleteAddressHandler
}

type createAddressRequest struct {
	UserID    int64  `json:"userId,string"` // not required for user, but required for admin
	Address   string `json:"address" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
}

type updateAddressRequest struct {
	ID        int64   `json:"id,string" binding:"required"`
	Address   *string `json:"address"`
	Phone     *string `json:"phone"`
	Recipient *string `json:"recipient"`
}

type deleteAddressRequest struct {
	UserID int64 `json:"userId,string"`
	ID     int64 `json:"id,string" binding:"required"`
}

func NewAddressHandler(ch *address.CreateAddressHandler, uh *address.UpdateAddressHandler, dh *address.DeleteAddressHandler) (*AddressHandler, error) {
	if ch == nil || uh == nil || dh == nil {
		return nil, errors.New("invalid address handler")
	}

	return &AddressHandler{createHandler: ch, updateHandler: uh, deleteHandler: dh}, nil
}

func (h *AddressHandler) CreateHandler(c *gin.Context) {
	var req createAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.createHandler.Handle(c.Request.Context(), &address.CreateAddressCommand{
		UserID:    req.UserID,
		Address:   req.Address,
		Phone:     req.Phone,
		Recipient: req.Recipient,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK[string](c, "address creation successfully")
}

func (h *AddressHandler) UpdateHandler(c *gin.Context) {
	var req updateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.updateHandler.Handle(c.Request.Context(), &address.UpdateAddressCommand{
		ID:        req.ID,
		Address:   req.Address,
		Phone:     req.Phone,
		Recipient: req.Recipient,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK[string](c, "address updating successfully")
}

func (h *AddressHandler) DeleteHandler(c *gin.Context) {
	var req deleteAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err := h.deleteHandler.Handle(c.Request.Context(), &address.DeleteAddressCommand{
		ID:     req.ID,
		UserID: req.UserID,
	})

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK[string](c, "address deleting successfully")
}
