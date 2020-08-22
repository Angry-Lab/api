package server

import (
	"github.com/Angry-Lab/api/pkg/segment"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	AuthHeader = "X-Auth-Token"
)

type Handler struct {
	Users    user.UseCase
	Segments segment.UseCase
}

func (h *Handler) WithAuth(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		t := ctx.Request().Header.Get(AuthHeader)
		if t == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "token required")
		}

		c := ctx.Request().Context()

		token, err := h.Users.SearchToken(c, t)
		if err != nil && err == user.ErrTokenNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error()).SetInternal(err)
		}

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
		}

		if token == nil {
			return echo.NewHTTPError(http.StatusNotFound)
		}

		ctx.Set(user.TokenKey, token)
		ctx.Set(user.ContextKey, token.User)

		return handlerFunc(ctx)
	}
}

func (h *Handler) WithAccess(roles ...user.Role) echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			u, ok := ctx.Get(user.ContextKey).(*user.User)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "no user")
			}

			for _, role := range roles {
				if role == u.Role {
					return handlerFunc(ctx)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "no access")
		}
	}
}
