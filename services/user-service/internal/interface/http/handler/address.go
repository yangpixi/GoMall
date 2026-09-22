package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/command/address"
	"github.com/yangpixi/GoMall/shared/errs"
	"github.com/yangpixi/GoMall/shared/response"
)

type AddressHandler struct {
	createHandler *address.CreateAddressHandler
	updateHandler *address.UpdateAddressHandler
}

type createAddressRequest struct {
	UserID    int64  `json:"userId" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
}

type updateAddressRequest struct {
	ID        int64   `json:"id" binding:"required"`
	Address   *string `json:"address"`
	Phone     *string `json:"phone"`
	Recipient *string `json:"recipient"`
}

func NewAddressHandler(ch *address.CreateAddressHandler, uh *address.UpdateAddressHandler) (*AddressHandler, error) {
	if ch == nil || uh == nil {
		return nil, errors.New("invalid address handler")
	}

	return &AddressHandler{createHandler: ch, updateHandler: uh}, nil
}

func (h *AddressHandler) CreateHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req createAddressRequest
	if err := c.ShouldBindJSON(req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	err := h.createHandler.Handle(c.Request.Context(), &address.CreateAddressCommand{
		UserID:    req.UserID,
		Address:   req.Address,
		Phone:     req.Phone,
		Recipient: req.Recipient,
	})

	if err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	response.OK[string](c, "address creation successfully")
}

func (h *AddressHandler) UpdateHandler(c *gin.Context) {
	var bizErr *errs.BusinessError
	var req updateAddressRequest
	if err := c.ShouldBindJSON(req); err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	err := h.updateHandler.Handle(c.Request.Context(), &address.UpdateAddressCommand{
		ID:        req.ID,
		Address:   req.Address,
		Phone:     req.Phone,
		Recipient: req.Recipient,
	})

	if err != nil {
		if errors.As(err, &bizErr) {
			_ = c.Error(bizErr)
			c.Abort()
			return
		}
	}

	response.OK[string](c, "address updating successfully")
}
