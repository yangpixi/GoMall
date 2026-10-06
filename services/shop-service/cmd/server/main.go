package main

import (
	"fmt"
	"log/slog"

	"github.com/yangpixi/GoMall/services/shop-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/shop-service/internal/config"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/shop-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/shop-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("shop", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic(fmt.Errorf("failed to init database connection: %w", err))
	}

	shopRepo := repository.NewShopRepo(db)
	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init id generator"))
	}

	createShopHandler, err := command.NewCreateShopHandler(shopRepo, generator)
	if err != nil {
		panic(fmt.Errorf("failed to init createShopHandler: %w", err))
	}

	shopHandler, err := handler.NewShopHandler(createShopHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init http shop handler: %w", err))
	}

	r := http.NewRouter(shopHandler, []byte(c.JWT.SecretKey))

	slog.Info("starting server", "port", c.Server.Port)

	if err = r.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to init http server"))
	}

	slog.Info("server starting successfully")
}
