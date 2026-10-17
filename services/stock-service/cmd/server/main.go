package main

import (
	"fmt"
	"log/slog"
	"net"

	stockv1 "github.com/yangpixi/GoMall/api/gen/go/stock/v1"
	"github.com/yangpixi/GoMall/services/stock-service/internal/application/command/stock"
	"github.com/yangpixi/GoMall/services/stock-service/internal/config"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/repository"
	grpciface "github.com/yangpixi/GoMall/services/stock-service/internal/interface/grpc"
	"github.com/yangpixi/GoMall/services/stock-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/stock-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
	"google.golang.org/grpc"
)

func main() {
	l := logger.New("stock", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	stockRepo := repository.NewStockRepo(db)
	reservationRepo := repository.NewReservationRepo(db)

	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init snowflake id generator: %w", err))
	}

	createHandler, err := stock.NewCreateSKUStockHandler(stockRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init create handler"))
	}

	httpStockHandler, err := handler.NewStockHandler(createHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init http stock handler: %w", err))
	}

	go func() {
		slog.Info("starting server", "port", c.Server.Port)

		router := http.NewRouter(httpStockHandler, []byte(c.JWT.SecretKey))
		if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
			panic(fmt.Errorf("failed to starting http server: %w", err))
		}
	}()

	// gRPC initialization
	reserveHandler, err := stock.NewReserveHandler(stockRepo, reservationRepo, generator)
	if err != nil {
		panic(fmt.Errorf("failed to init reserve handler"))
	}

	releaseHandler, err := stock.NewReleaseHandler(reservationRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init release handler"))
	}

	confirmHandler, err := stock.NewConfirmHandler(reservationRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init confirm handler"))
	}

	stockServer := grpciface.NewStockServer(reserveHandler, releaseHandler, confirmHandler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Server.GRPCPort))
	if err != nil {
		panic(fmt.Errorf("failed to listen port: %d, error: %w", c.Server.GRPCPort, err))
	}

	srv := grpc.NewServer()
	stockv1.RegisterStockServiceServer(srv, stockServer)

	slog.Info("starting gRPC server", "port", c.Server.GRPCPort)
	if err = srv.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to start gRPC server"))
	}
}
