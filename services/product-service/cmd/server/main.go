package main

import (
	"fmt"
	"log/slog"

	shopv1 "github.com/yangpixi/GoMall/api/gen/go/shop/v1"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/command/product"
	"github.com/yangpixi/GoMall/services/product-service/internal/application/command/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/config"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/rpc"
	"github.com/yangpixi/GoMall/services/product-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/product-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init id generator"))
	}

	conn, err := grpc.NewClient("localhost:9094",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to shop gRPC server: %w", err))
	}

	cli := rpc.NewShopClient(shopv1.NewShopServiceClient(conn))

	appProductHandler, err := product.NewCreateProductHandler(productRepo, generator, cli)
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
