package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/payment-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(notificationHandler *handler.NotificationHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	payment := r.Group("/api/v1/payment")
	payment.POST("/notify/mock", middleware.RequireJWT(secretKey), notificationHandler.MockChannelHandler)
	payment.POST("/notify/wechat", middleware.RequireJWT(secretKey), notificationHandler.WechatChannelHandler)
	payment.POST("/notify/alipay", middleware.RequireJWT(secretKey), notificationHandler.AlipayChannelHandler)

	return r
}
