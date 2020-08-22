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

	ok, err := h.Users.ComparePassword(u, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password").SetInternal(err)
	}

	t, err := h.Users.CreateAuthToken(ctx.Request().Context(), u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	u.Password = ""
	t.User = u

	return ctx.JSON(http.StatusOK, t)
}

type authRequest struct {
	Login    string
	Password string
}
