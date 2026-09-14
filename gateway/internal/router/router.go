package router

import (
	"net/http"

	"github.com/yangpixi/GoMall/gateway/internal/proxy"
	"github.com/yangpixi/GoMall/gateway/internal/router/middleware"
)

// New return an http router that defines all proxy rules.
func New(secretKey []byte) *http.ServeMux {
	mux := http.NewServeMux()
	requiredJWT := middleware.RequiredJWT(secretKey)

	mux.Handle("/api/v1/auth/", proxy.New("http://localhost:8081"))
	mux.Handle("/api/v1/user/", requiredJWT(proxy.New("http://localhost:8082")))

	return mux
}
