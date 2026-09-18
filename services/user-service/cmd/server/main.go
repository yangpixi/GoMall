package main

import (
	"fmt"
	"log/slog"

	"github.com/yangpixi/GoMall/services/user-service/internal/application/command"
	"github.com/yangpixi/GoMall/services/user-service/internal/config"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/services/user-service/internal/interface/http"
	httpHandler "github.com/yangpixi/GoMall/services/user-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("user", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.yaml")
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic(fmt.Errorf("failed to init database connection: %w", err))
	}

	profileRepo := repository.NewProfileRepo(db)

	handler, err := command.NewCreateProfileHandler(profileRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init profile handler: %w", err))
	}

	profileHandler, err := httpHandler.NewProfileHandler(handler)
	if err != nil {
		panic(fmt.Errorf("failed to init profile http handler: %w", err))
	}

	router := http.NewRouter(profileHandler, []byte(c.JWT.SecretKey))
	slog.Info("starting server", "port", c.Server.Port)
	if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to serve http server on the specific port: %d, error: %w", c.Server.Port, err))
	}

}
