package main

import (
	"fmt"
	"log/slog"

	"github.com/yangpixi/GoMall/services/stock-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/stock-service/internal/config"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/stock-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/stock-service/internal/interface/http"
	"github.com/yangpixi/GoMall/services/stock-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
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

	reserveHandler, err := command.NewReserveHandler(stockRepo, reservationRepo, generator)
	httpStockHandler, err := handler.NewStockHandler(reserveHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init http stock handler: %w", err))
	}

	slog.Info("starting server", "port", c.Server.Port)

	router := http.NewRouter(httpStockHandler, []byte(c.JWT.SecretKey))
	if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to starting http server: %w", err))
	}

}
