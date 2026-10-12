package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/stock-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(stockHandler *handler.StockHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	shop := r.Group("/api/v1/stock")
	shop.POST("/reserve", middleware.RequireJWT(secretKey), stockHandler.ReserveHandler)

	return r
}
