package handler

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, shortenHandler ShortenHandler) {
	e.POST("/shorten", shortenHandler.Shorten)
	e.GET("/:short_id", shortenHandler.Resolve)

	e.Use(echoprometheus.NewMiddleware("sdalm"))
	e.GET("/metrics", echoprometheus.NewHandler())
}
