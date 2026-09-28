package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/product-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(productHandler *handler.ProductHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	product := r.Group("/api/v1/product")
	product.POST("/create", middleware.RequireJWT(secretKey), productHandler.CreateProductHandler)

	sku := r.Group("/api/v1/sku")
	sku.POST("/create", middleware.RequireJWT(secretKey), productHandler.CreateSKUHandler)

	return r
}
