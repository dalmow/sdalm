package http

import "github.com/labstack/echo/v4"

func RegisterRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})
}
