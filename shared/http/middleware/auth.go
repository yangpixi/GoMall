package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yangpixi/GoMall/shared/http/id"
)

type claims struct {
	RoleIDs []int64 `json:"role_ids"`
	Kind    string  `json:"kind"`
	jwt.RegisteredClaims
}

// RequireJWT parsing the jwt and inject it into context
func RequireJWT(secretKey []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fields := strings.Fields(ctx.GetHeader("Authorization"))

		if len(fields) != 2 || !strings.EqualFold(fields[0], "Bearer") {
			unauthorize(ctx)
			return
		}

		token := fields[1]
		if token == "" {
			unauthorize(ctx)
			return
		}

		res, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (any, error) {
			return secretKey, nil
		},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			jwt.WithLeeway(5*time.Second),
		)

		if err != nil || !res.Valid {
			slog.Info("failed to parse token", "error", err)
			unauthorize(ctx)
			return
		}

		if c, ok := res.Claims.(*claims); ok && c.Kind == "access" {
			i, err := strconv.ParseInt(c.Subject, 10, 64)
			if err != nil {
				slog.Info("failed to parse userID", "id", c.Subject, "error", err)
				unauthorize(ctx)
				return
			}

			ctxWithID := id.WithUserID(ctx.Request.Context(), i)
			ctx.Request = ctx.Request.WithContext(ctxWithID)
		} else {
			unauthorize(ctx)
			return
		}

		// pass
		ctx.Next()
	}
}

func unauthorize(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  "missing token",
	})
}
