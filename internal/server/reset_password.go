package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) ResetPassword(ctx echo.Context) error {
	return ctx.NoContent(http.StatusOK)
}
