package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/shared/errs"
)

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
}

func OK[T any](c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response[T]{
		Code: 0,
		Data: data,
	})
}

func Fail(c *gin.Context, err *errs.BusinessError) {
	c.JSON(http.StatusOK, Response[struct{}]{
		Code: err.Code,
		Msg:  err.Msg,
	})
}
