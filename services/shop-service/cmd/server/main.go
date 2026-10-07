package main

import (
	"fmt"
	"log/slog"
	"net"

	shopv1 "github.com/yangpixi/GoMall/api/gen/go/shop/v1"
	"github.com/yangpixi/GoMall/services/shop-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/shop-service/internal/application/query/shop"
	"github.com/yangpixi/GoMall/services/shop-service/internal/config"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/reader"
	"github.com/yangpixi/GoMall/services/shop-service/internal/infrastructure/persistence/postgres/repository"
	grpciface "github.com/yangpixi/GoMall/services/shop-service/internal/interface/grpc"
	"github.com/yangpixi/GoMall/services/shop-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/shop-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
	"google.golang.org/grpc"
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

	go func() {
		r := http.NewRouter(shopHandler, []byte(c.JWT.SecretKey))

		slog.Info("starting server", "port", c.Server.Port)

		if err = r.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
			panic(fmt.Errorf("failed to start http server"))
		}
	}()

	// gRPC initialization

	sr := reader.NewShopReader(db)
	shopIDHandler, err := shop.NewGetShopIDHandler(sr)
	if err != nil {
		panic(fmt.Errorf("failed to init shop query handler: %w", err))
	}

	shopServer := grpciface.NewShopServer(shopIDHandler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Server.GRPCPort))
	if err != nil {
		panic(fmt.Errorf("failed to listen port: %d, error: %w", c.Server.GRPCPort, err))
	}

	srv := grpc.NewServer()
	shopv1.RegisterShopServiceServer(srv, shopServer)

	slog.Info("starting gRPC server", "port", c.Server.GRPCPort)
	if err = srv.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to start gRPC server"))
	}
}
