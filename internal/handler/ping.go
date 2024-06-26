package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /ping
func (h *Handler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
