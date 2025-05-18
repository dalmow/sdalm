package handler

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, shortenHandler ShortenHandler) {
	e.POST("/shorten", shortenHandler.Shorten)
}
