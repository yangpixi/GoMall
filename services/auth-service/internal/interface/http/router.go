package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/auth-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/auth-service/internal/interface/http/middleware"
)

func NewRouter(handler *handler.AuthHandler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.ErrorHandler())

	auth := r.Group("/api/v1/auth")
	auth.POST("/login", handler.LoginHandler)
	auth.POST("/register", handler.RegisterHandler)
	auth.POST("/refresh", handler.RefreshHandler)

	return r
}
