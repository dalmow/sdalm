package http

import (
	"net/http"

	"github.com/dalmow/sdalm/internal/data"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *data.Database) {
	e.GET("/healthcheck", func(c echo.Context) error {
		if err := db.Pool.Ping(c.Request().Context()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	})
}
