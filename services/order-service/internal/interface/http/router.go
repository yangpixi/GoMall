package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/order-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(orderHandler *handler.OrderHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	product := r.Group("/api/v1/order")
	product.POST("/create", middleware.RequireJWT(secretKey), orderHandler.CreateHandler)

	return r
}
