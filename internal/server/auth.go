package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) Auth(ctx echo.Context) error {
	req := &authRequest{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	u, err := h.Users.SearchByLogin(ctx.Request().Context(), req.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	if u.Password != req.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	t, err := h.Users.CreateAuthToken(ctx.Request().Context(), u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return ctx.JSON(http.StatusOK, t)
}

type authRequest struct {
	Login    string
	Password string
}
