package main

import (
	"fmt"
	"log/slog"
	"net"

	productv1 "github.com/yangpixi/GoMall/api/gen/go/product/v1"
	shopv1 "github.com/yangpixi/GoMall/api/gen/go/shop/v1"
	productcommand "github.com/yangpixi/GoMall/services/product-service/internal/application/command/product"
	skucommand "github.com/yangpixi/GoMall/services/product-service/internal/application/command/sku"
	skuquery "github.com/yangpixi/GoMall/services/product-service/internal/application/query/sku"
	"github.com/yangpixi/GoMall/services/product-service/internal/config"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/reader"
	"github.com/yangpixi/GoMall/services/product-service/internal/infrastructure/rpc"
	grpciface "github.com/yangpixi/GoMall/services/product-service/internal/interface/grpc"
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

	appProductHandler, err := productcommand.NewCreateProductHandler(productRepo, generator, cli)
	if err != nil {
		panic(fmt.Errorf("failed to init application product handler: %w", err))
	}

	appSKUHandler, err := skucommand.NewCreateSKUHandler(skuRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init application sku handler: %w", err))
	}

	httpHandler, err := handler.NewProductHandler(appProductHandler, appSKUHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init http product handler: %w", err))
	}

	router := http.NewRouter(httpHandler, []byte(c.JWT.SecretKey))

	go func() {
		slog.Info("starting http server", "port", c.Server.Port)
		if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
			panic(fmt.Errorf("failed to run http server: %w", err))
		}
	}()

	// grpc initialization

	skuReader := reader.NewSKUReader(db)
	querySKUHandler := skuquery.NewGetSKUPriceHandler(skuReader)
	productServer := grpciface.NewProductServer(querySKUHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Server.GRPCPort))
	if err != nil {
		panic(fmt.Errorf("failed to listen port: %d, error: %w", c.Server.GRPCPort, err))
	}

	srv := grpc.NewServer()
	productv1.RegisterSKUServiceServer(srv, productServer)

	slog.Info("starting gRPC server", "port", c.Server.GRPCPort)
	if err = srv.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to start gRPC server"))
	}
}
