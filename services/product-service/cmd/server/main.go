package main

import (
	"fmt"
	"log/slog"

	"github.com/yangpixi/GoMall/services/product-service/internal/application/command/product"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/command/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/config"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/product-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/product-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("product", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic(fmt.Errorf("failed to init database connection: %w", err))
	}

	productRepo := repository.NewProductRepo(db)
	skuRepo := repository.NewSKURepo(db)

	appProductHandler, err := product.NewCreateProductHandler(productRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init application product handler: %w", err))
	}

	appSKUHandler, err := sku.NewCreateSKUHandler(skuRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init application sku handler: %w", err))
	}

	httpHandler, err := handler.NewProductHandler(appProductHandler, appSKUHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init http product handler: %w", err))
	}

	router := http.NewRouter(httpHandler, []byte(c.JWT.SecretKey))

	slog.Info("starting server", "port", c.Server.Port)

	if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to run http server: %w", err))
	}

}
