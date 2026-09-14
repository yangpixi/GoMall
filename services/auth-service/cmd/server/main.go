package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/yangpixi/GoMall/auth-service/internal/application/command"
	"github.com/yangpixi/GoMall/auth-service/internal/config"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/jwt"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/auth-service/internal/interface/http"
	"github.com/yangpixi/GoMall/auth-service/internal/interface/http/handler"
	"github.com/yangpixi/GoMall/shared/logger"
)

func main() {
	l := logger.New("auth", slog.LevelInfo)
	slog.SetDefault(l)

	c, err := config.Load("./config/config.example.yaml")

	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	accountRepo := repository.NewAccountRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	issuer, err := jwt.NewJWT(
		c.JWT.SecretKey,
		time.Duration(c.JWT.Expiration)*time.Second,
		time.Duration(c.JWT.RefreshExpiration)*time.Second,
	)
	if err != nil {
		panic(fmt.Errorf("failed to init jwt issuer: %w", err))
	}

	loginHandler, err := command.NewLoginHandler(accountRepo, issuer)
	if err != nil {
		panic(fmt.Errorf("failed to init login handler: %w", err))
	}

	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init snowflake generator: %w", err))
	}

	registerHandler, err := command.NewRegisterHandler(accountRepo, roleRepo, generator)
	if err != nil {
		panic(fmt.Errorf("failed to init register handler: %w", err))
	}

	refreshHandler, err := command.NewRefreshHandler(issuer)
	if err != nil {
		panic(fmt.Errorf("failed to init refresh handler: %w", err))
	}

	authHandler, err := handler.NewAuthHandler(loginHandler, registerHandler, refreshHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init authHandler: %w", err))
	}

	r := http.NewRouter(authHandler)

	if err = r.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to run http server on specific port: %w", err))
	}

	slog.Info("service starts successfully", "port", c.Server.Port)
}
