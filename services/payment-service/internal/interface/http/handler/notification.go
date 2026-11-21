package handler

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/payment-service/internal/application/command/payment"
	"github.com/yangpixi/GoMall/shared/response"
)

type NotificationHandler struct {
	callbackHandler *payment.CallbackHandler
}

func NewNotificationHandler(cbh *payment.CallbackHandler) *NotificationHandler {
	return &NotificationHandler{callbackHandler: cbh}
}

type mockNotification struct {
	TradeNo      string
	PaymentID    int64
	Status       string
	CallbackTime string
}

func (h *NotificationHandler) MockChannelHandler(c *gin.Context) {
	req := &mockNotification{}
	if err := c.ShouldBindJSON(req); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	t, err := time.Parse(time.RFC3339, req.CallbackTime)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	err = h.callbackHandler.Handle(c.Request.Context(), &payment.CallbackCommand{
		TradeNo:      req.TradeNo,
		PaymentID:    req.PaymentID,
		Status:       req.Status,
		CallbackTime: t,
	})
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	response.OK(c, "success")
}

func (h *NotificationHandler) WechatChannelHandler(c *gin.Context) {

}

func (h *NotificationHandler) AlipayChannelHandler(c *gin.Context) {

}
