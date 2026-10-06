package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/shop-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(shopHandler *handler.ShopHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	shop := r.Group("/api/v1/shop")
	shop.POST("/create", middleware.RequireJWT(secretKey), shopHandler.CreateHandler)

	return r
}
