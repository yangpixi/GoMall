package middleware

import (
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type claims struct {
	RoleIDs []int64 `json:"role_ids"`
	Kind    string  `json:"kind"`
	jwt.RegisteredClaims
}

// RequiredJWT an http middleware that checks user's Bearer token
func RequiredJWT(secretKey []byte) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			token := strings.Fields(req.Header.Get("Authorization"))

			if len(token) != 2 || !strings.EqualFold(token[0], "Bearer") {
				unauthorized(w)
				return
			}

			res, err := jwt.ParseWithClaims(token[1], &claims{}, func(token *jwt.Token) (any, error) {
				return secretKey, nil
			},
				jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
				jwt.WithLeeway(5*time.Second),
			)

			if err != nil || !res.Valid {
				slog.Error("failed to parse token")
				unauthorized(w)
				return
			}

			if c, ok := res.Claims.(*claims); !ok || c.Kind != "access" {
				unauthorized(w)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"code":"401","message":"please login or relogin"}`))
}
