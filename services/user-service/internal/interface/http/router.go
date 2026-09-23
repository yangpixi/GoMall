package http

import (
	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/services/user-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/http/middleware"
)

func NewRouter(profileHandler *handler.ProfileHandler, addressHandler *handler.AddressHandler, secretKey []byte) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())
	r.Use(middleware.ErrorHandler())

	user := r.Group("/api/v1/user")
	user.POST("/profile/create", middleware.RequireJWT(secretKey), profileHandler.CreateHandler)
	user.POST("/profile/update", middleware.RequireJWT(secretKey), profileHandler.UpdateHandler)
	user.GET("/profile", middleware.RequireJWT(secretKey), profileHandler.DetailHandler)

	user.POST("/address/create", middleware.RequireJWT(secretKey), addressHandler.CreateHandler)
	user.POST("/address/update", middleware.RequireJWT(secretKey), addressHandler.UpdateHandler)
	user.POST("/address/delete", middleware.RequireJWT(secretKey), addressHandler.DeleteHandler)

	return r
}
