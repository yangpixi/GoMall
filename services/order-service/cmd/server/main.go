package main

import (
	"fmt"
	"log/slog"

	productv1 "github.com/yangpixi/GoMall/api/gen/go/product/v1"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/order-service/internal/config"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/rpc"
	"github.com/yangpixi/GoMall/services/order-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/order-service/internal/interface/http/handler"
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
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	orderRepo := repository.NewOrderRepo(db)
	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init id generator: %w", err))
	}

	conn, err := grpc.NewClient("localhost:9093", // 9093 is product gRPC server port
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to init gRPC connection"))
	}

	productClient := rpc.NewProductClient(productv1.NewSKUServiceClient(conn))

	createHandler := command.NewCreateHandler(orderRepo, generator, productClient)

	orderHandler := handler.NewOrderHandler(createHandler)

	router := http.NewRouter(orderHandler, []byte(c.JWT.SecretKey))

	slog.Info("starting http server", "port", c.Server.Port)
	if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to run http server: %w", err))
	}
}
