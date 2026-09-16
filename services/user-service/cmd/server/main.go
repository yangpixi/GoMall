package main

import (
	"fmt"
	"log/slog"

	"github.com/yangpixi/GoMall/services/user-service/internal/config"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("user", slog.LevelInfo)
	slog.SetDefault(l)

	_, err := config.Load("../config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

}
