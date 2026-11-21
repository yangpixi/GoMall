package main

import (
	"fmt"
	"log/slog"
	"net"

	orderv1 "github.com/yangpixi/GoMall/api/gen/go/order/v1"
	paymentv1 "github.com/yangpixi/GoMall/api/gen/go/payment/v1"
	paymentcommand "github.com/yangpixi/GoMall/services/payment-service/internal/application/command/payment"
	"github.com/yangpixi/GoMall/services/payment-service/internal/config"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/id"
	paymentinfra "github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/payment"
	_ "github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/payment/mock"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/payment-service/internal/infrastructure/rpc"
	grpciface "github.com/yangpixi/GoMall/services/payment-service/internal/interface/grpc"
	"github.com/yangpixi/GoMall/services/payment-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/payment-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	l := logger.New("payment", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to init config: %w", err))
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic(fmt.Errorf("failed to init database connection: %w", err))
	}

	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init snowflake id generator: %w", err))
	}

	conn, err := grpc.NewClient("localhost:9096", // 9096 is order gRPC server port
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to connect to order gRPC server: %w", err))
	}

	orderClient := rpc.NewOrderClient(orderv1.NewOrderServiceClient(conn))

	paymentRepo := repository.NewPaymentRepo(db)
	callbackHandler := paymentcommand.NewCallbackHandler(paymentRepo)

	notificationHandler := handler.NewNotificationHandler(callbackHandler)
	router := http.NewRouter(notificationHandler, []byte(c.JWT.SecretKey))

	go func() {
		slog.Info("starting http server", "port", c.Server.Port)
		if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
			panic(fmt.Errorf("failed to run http server: %w", err))
		}
	}()

	// gRPC
	registry, err := paymentinfra.NewRegistry(c)
	if err != nil {
		panic(fmt.Errorf("failed to init registry: %w", err))
	}

	createPaymentHandler := paymentcommand.NewCreatePaymentHandler(paymentRepo, generator, registry, orderClient)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Server.GRPCPort))
	if err != nil {
		panic(fmt.Errorf("failed to create listening on port: %d", c.Server.GRPCPort))
	}

	srv := grpc.NewServer()
	paymentServer := grpciface.NewPaymentServer(createPaymentHandler)
	paymentv1.RegisterPaymentServiceServer(srv, paymentServer)

	if err = srv.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to listen port: %w", err))
	}
}
