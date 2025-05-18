package main

import (
	"github.com/dalmow/sdalm/internal/config"
	"github.com/dalmow/sdalm/internal/data"
	"github.com/dalmow/sdalm/internal/handler"
	"github.com/dalmow/sdalm/internal/logger"
	"github.com/dalmow/sdalm/internal/short"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			data.NewDatabaseConnection,
			logger.NewLogger,
			short.NewShortsRepository,
			short.NewShortenUseCase,
			handler.NewShortenHandler,
			handler.NewHttpServer),
		fx.Invoke(
			data.RunMigrations,
			data.CloseDatabaseConnection,
			handler.RegisterRoutes,
			handler.StartServer))

	app.Run()
}
