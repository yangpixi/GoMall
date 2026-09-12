package main

import (
	"fmt"
	"time"

	"github.com/yangpixi/GoMall/auth-service/internal/application/command"
	"github.com/yangpixi/GoMall/auth-service/internal/config"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/jwt"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/auth-service/internal/infrastructure/persistence/postgres/repository"
	"github.com/yangpixi/GoMall/auth-service/internal/interface/http"
	"github.com/yangpixi/GoMall/auth-service/internal/interface/http/handler"
)

func main() {
	c, err := config.Load("./config/config.yaml")

	if err != nil {
		panic(err)
	}

	db, err := postgres.NewDB(c)
	if err != nil {
		panic("failed to connect to database")
	}

	accountRepo := repository.NewAccountRepo(db)
	roleRepo := repository.NewRoleRepo(db)
	issuer, err := jwt.NewJWT(
		c.JWT.SecretKey,
		time.Duration(c.JWT.Expiration)*time.Second,
		time.Duration(c.JWT.RefreshExpiration)*time.Second,
	)
	if err != nil {
		panic("failed to init jwt issuer")
	}

	loginHandler, err := command.NewLoginHandler(accountRepo, issuer)
	if err != nil {
		panic("failed to init login handler")
	}

	generator, err := id.NewGenerator(1)
	if err != nil {
		panic("failed to init snowflake generator")
	}

	registerHandler, err := command.NewRegisterHandler(accountRepo, roleRepo, generator)
	if err != nil {
		panic("failed to init register handler")
	}

	authHandler, err := handler.NewAuthHandler(loginHandler, registerHandler)
	if err != nil {
		panic("failed to init authHandler")
	}

	r := http.NewRouter(authHandler)

	if err = r.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic("failed to run http server on specific port")
	}
}
