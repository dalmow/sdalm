package handler

import (
	"errors"
	"net/http"

	"github.com/dalmow/sdalm/internal/short"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ShortenHandler interface {
	Shorten(ctx echo.Context) error
	Resolve(ctx echo.Context) error
}

type shortenHandler struct {
	usecase short.ShortenUseCase
	logger  *zap.Logger
}

func NewShortenHandler(usecase short.ShortenUseCase, logger *zap.Logger) ShortenHandler {
	return &shortenHandler{
		usecase: usecase,
		logger:  logger,
	}
}

func (h *shortenHandler) Shorten(ctx echo.Context) error {
	type request struct {
		URL string `json:"url" validate:"required,url"`
	}
	req := new(request)
	if err := ctx.Bind(&req); err != nil {
		h.logger.Warn("invalid request", zap.Error(err))
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if err := ctx.Validate(req); err != nil {
		h.logger.Warn("invalid request", zap.Error(err))
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{"error": err.Error()})
	}

	shortened, err := h.usecase.ShortenURL(ctx.Request().Context(), req.URL)
	if err != nil {
		h.logger.Error("shorten use case error", zap.Error(err))
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "could not shorten URL"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"short": shortened,
	})
}

func (h *shortenHandler) Resolve(ctx echo.Context) error {
	identifier := ctx.Param("short_id")
	originalUrl, err := h.usecase.Resolve(ctx.Request().Context(), identifier)

	if err != nil {
		if errors.Is(err, short.ErrNotFound) {
			return ctx.NoContent(http.StatusNotFound)
		}
		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.Redirect(http.StatusFound, originalUrl)
}
