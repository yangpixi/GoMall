package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(h *handler.ProfileHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	user := r.Group("/api/v1/user")
	user.POST("/profile/create", middleware.RequireJWT(secretKey), h.Handler)

	return r
}
