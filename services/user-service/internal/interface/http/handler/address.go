package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/command"
	"github.com/yangpixi/GoMall/shared/errs"
	"github.com/yangpixi/GoMall/shared/response"
)

type AddressHandler struct {
	handler *command.CreateAddressHandler
}

type createAddressRequest struct {
	UserID    int64
	Address   string
	Phone     string
	Recipient string
}

func NewAddressHandler(h *command.CreateAddressHandler) (*AddressHandler, error) {
	if h == nil {
		return nil, errors.New("invalid address handler")
	}

	return &AddressHandler{handler: h}, nil
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

	err := h.handler.Handle(c.Request.Context(), &command.CreateAddressCommand{
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
