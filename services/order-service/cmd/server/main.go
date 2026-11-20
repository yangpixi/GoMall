package main

import (
	"fmt"
	"log/slog"
	"net"

	orderv1 "github.com/yangpixi/GoMall/api/gen/go/order/v1"
	productv1 "github.com/yangpixi/GoMall/api/gen/go/product/v1"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/order-service/internal/application/query"
	"github.com/yangpixi/GoMall/services/order-service/internal/config"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/reader"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/order-service/internal/infrastructure/rpc"
	grpciface "github.com/yangpixi/GoMall/services/order-service/internal/interface/grpc"
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

	go func() {
		slog.Info("starting http server", "port", c.Server.Port)
		if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
			panic(fmt.Errorf("failed to run http server: %w", err))
		}
	}()

	// gRPC

	orderReader := reader.NewOrderReader(db)
	orderQueryHandler := query.NewOrderQueryHandler(orderReader)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Server.GRPCPort))
	if err != nil {
		panic(fmt.Errorf("failed to create listening on port: %d", c.Server.GRPCPort))
	}

	srv := grpc.NewServer()
	orderServer := grpciface.NewOrderServer(createHandler, orderQueryHandler)
	orderv1.RegisterOrderServiceServer(srv, orderServer)

	if err = srv.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to listen port: %w", err))
	}
}
