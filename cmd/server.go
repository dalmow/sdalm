package main

import (
	"github.com/dalmow/sdalm/internal/config"
	"github.com/dalmow/sdalm/internal/http"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			http.NewHttpServer,
			config.LoadConfig),
		fx.Invoke(
			http.RegisterRoutes,
			http.StartServer))

	app.Run()
}
