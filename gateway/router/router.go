package router

import (
	"net/http"

	"github.com/yangpixi/GoMall/gateway/proxy"
	"github.com/yangpixi/GoMall/gateway/router/middleware"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	requiredJWT := middleware.RequiredJWT([]byte("123132"))

	mux.Handle("/api/v1/auth/", proxy.New("http://localhost:8081"))
	mux.Handle("/api/v1/user/", requiredJWT(proxy.New("http://localhost:8082")))

	return mux
}
