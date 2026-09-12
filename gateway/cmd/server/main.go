package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/yangpixi/GoMall/gateway/router"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	// setting the logger
	l := logger.New("gateway", slog.LevelInfo)
	slog.SetDefault(l)

	slog.Info("service starting successfully", "port", "8080")

	mux := router.New()

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error("http server starting failed", "error", err)
		panic(err)
	}
}
