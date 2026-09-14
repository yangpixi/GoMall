package proxy

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// New return a reverse proxy instance
func New(path string) http.Handler {
	u, err := url.Parse(path)
	if err != nil {
		slog.Error("url parsing failed")
	}

	proxy := httputil.NewSingleHostReverseProxy(u)

	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
		slog.Error("proxy failed", "url", path, "error", err)

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadGateway)
		_, _ = writer.Write([]byte(
			`{"code":"502","message":"Service is temporarily unavailable"}`,
		))
	}

	return proxy
}
