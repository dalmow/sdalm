package main

import (
	"github.com/dalmow/sdalm/internal/config"
	"github.com/dalmow/sdalm/internal/data"
	"github.com/dalmow/sdalm/internal/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.LoadConfig,
			data.NewDatabaseConnection,
			http.NewHttpServer),
		fx.Invoke(
			data.RunMigrations,
			data.CloseDatabaseConnection,
			http.RegisterRoutes,
			http.StartServer))

	app.Run()
}
