package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/yangpixi/GoMall/gateway/internal/config"
	"github.com/yangpixi/GoMall/gateway/internal/router"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	// setting the logger
	l := logger.New("gateway", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	mux := router.New([]byte(c.JWT.SecretKey))

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", c.Server.Port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	slog.Info("starting server", "port", c.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		panic(fmt.Errorf("http server starting failed: %w", err))
	}

}
