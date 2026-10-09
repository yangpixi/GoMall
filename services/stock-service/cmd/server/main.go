package main

import (
	"log/slog"

	"github.com/yangpixi/GoMall/services/stock-service/internal/config"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("stock", slog.LevelInfo)
	slog.SetDefault(l)

	config.Load("./config/config.yaml")
}
