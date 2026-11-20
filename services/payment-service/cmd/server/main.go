package main

import (
	"log/slog"

	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("payment", slog.LevelInfo)
	slog.SetDefault(l)

}
