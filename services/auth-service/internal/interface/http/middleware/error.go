package middleware

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangpixi/GoMall/shared/errs"
	"github.com/yangpixi/GoMall/shared/response"
)

// ErrorHandler process business error, which extracts error code and
// converts it to http status code
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// ignore when there is no error
		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		err := c.Errors.Last().Err
		var bizErr *errs.BusinessError
		if errors.As(err, &bizErr) {
			response.Fail(c, bizErr, statusFromCode(bizErr.Code))
			return
		}

		slog.ErrorContext(c.Request.Context(), "failed to process request",
			"path", c.Request.URL.Path,
			"error", err,
		)

		response.Fail(c, errs.New(500, "internal server error"), http.StatusInternalServerError)
	}
}

// refer to domain/*/errors
func statusFromCode(code int) int {
	switch code {
	case 10001:
		return http.StatusForbidden
	case 10002, 10004, 10005, 10006:
		return http.StatusBadRequest
	case 10003:
		return http.StatusUnauthorized
	default:
		return http.StatusBadRequest
	}
}
