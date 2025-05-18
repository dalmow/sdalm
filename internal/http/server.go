package http

import (
	"context"
	"fmt"

	"github.com/dalmow/sdalm/internal/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewHttpServer() *echo.Echo {
	return echo.New()
}

func StartServer(lc fx.Lifecycle, e *echo.Echo, c *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(fmt.Sprintf(":%v", c.AppPort)); err != nil {
					e.Logger.Fatal(fmt.Sprintf("Error when server start: %v", err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}
