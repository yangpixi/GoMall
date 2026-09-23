package main

import (
	"fmt"
	"log/slog"

	"github.com/yangpixi/GoMall/services/user-service/internal/application/command/address"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/command/profile"
	"github.com/yangpixi/GoMall/services/user-service/internal/application/query"
	"github.com/yangpixi/GoMall/services/user-service/internal/config"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/id"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres"
	"github.com/yangpixi/GoMall/services/user-service/internal/infrastructure/persistence/postgres/reader"
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
	addressRepo := repository.NewAddressRepo(db)

	generator, err := id.NewGenerator(1)
	if err != nil {
		panic(fmt.Errorf("failed to init id generator: %w", err))
	}

	appProfileCreateHandler, err := profile.NewCreateProfileHandler(profileRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init profile handler: %w", err))
	}

	appProfileUpdateHandler, err := profile.NewUpdateProfileHandler(profileRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init profile handler: %w", err))
	}

	profileReader := reader.NewProfileReader(db)
	getProfileHandler, err := query.NewGetProfileHandler(profileReader)
	if err != nil {
		panic(fmt.Errorf("failed to init profile handler: %w", err))
	}

	profileHandler, err := httpHandler.NewProfileHandler(appProfileCreateHandler, appProfileUpdateHandler, getProfileHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init profile http handler: %w", err))
	}

	appAddressCreateHandler, err := address.NewCreateAddressHandler(addressRepo, generator)
	if err != nil {
		panic(fmt.Errorf("failed to init address handler: %w", err))
	}

	appAddressUpdateHandler, err := address.NewUpdateAddressHandler(addressRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init address handler: %w", err))
	}

	appAddressDeleteHandler, err := address.NewDeleteAddressHandler(addressRepo)
	if err != nil {
		panic(fmt.Errorf("failed to init address handler: %w", err))
	}

	addressHandler, err := httpHandler.NewAddressHandler(appAddressCreateHandler, appAddressUpdateHandler, appAddressDeleteHandler)
	if err != nil {
		panic(fmt.Errorf("failed to init addressHandler"))
	}

	router := http.NewRouter(profileHandler, addressHandler, []byte(c.JWT.SecretKey))
	slog.Info("starting server", "port", c.Server.Port)
	if err = router.Run(fmt.Sprintf(":%d", c.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to serve http server on the specific port: %d, error: %w", c.Server.Port, err))
	}

}
