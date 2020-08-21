package server

import (
	"context"
	"fmt"
	"github.com/Angry-Lab/api/internal/env"
	"github.com/Angry-Lab/api/pkg/entity"
	"github.com/Angry-Lab/api/pkg/mockup"
	"github.com/Angry-Lab/api/pkg/token"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/partyzanex/testutils"
	"net/http"
	"time"
)

type Handler struct {
	Tokens token.Repository
	Users  user.Repository
	Count  int64
}

type Request struct {
	Login    string
	Password string
}

func (h *Handler) Auth(ctx echo.Context) error {
	req := &Request{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	u, err := h.Users.GetOneByLogin(ctx.Request().Context(), req.Login)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	if u.Password != req.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid password")
	}

	t := &entity.Token{
		Value:     testutils.RandomString(64),
		UserID:    u.ID,
		DTCreated: time.Now().UTC(),
		DTUsed:    nil,
		DTExpired: time.Now().Add(time.Hour * 24).UTC(),
	}

	_ = h.Tokens.Set(ctx.Request().Context(), t)

	return ctx.JSON(http.StatusOK, t)
}

func (h *Handler) Register(ctx echo.Context) error {
	req := &entity.User{}

	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error()).SetInternal(err)
	}

	req.ID = h.Count
	h.Count++

	_ = h.Users.Update(ctx.Request().Context(), req)

	t := &entity.Token{
		Value:     testutils.RandomString(64),
		UserID:    req.ID,
		DTCreated: time.Now().UTC(),
		DTUsed:    nil,
		DTExpired: time.Now().Add(time.Hour * 24).UTC(),
	}

	_ = h.Tokens.Set(ctx.Request().Context(), t)

	return ctx.JSON(http.StatusOK, t)
}

func (h *Handler) ResetPassword(ctx echo.Context) error {
	return nil
}

func Run(config *env.Config) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodPost, http.MethodGet, http.MethodOptions},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodPost, http.MethodGet, http.MethodOptions},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"},
		AllowCredentials: true,
	}))

	h := &Handler{
		Tokens: mockup.Tokens(),
		Users:  mockup.Users(),
		Count:  2,
	}

	ctx := context.Background()

	_ = h.Users.Update(ctx, &entity.User{
		ID:           1,
		Name:         "Test User",
		Login:        "login@email.com",
		Password:     "qwerty",
		Status:       1,
		Role:         "user",
		DTCreated:    time.Now().UTC(),
		DTUpdated:    time.Now().UTC(),
		DTLastLogged: nil,
	})

	e.POST("/auth", h.Auth)
	e.POST("/register", h.Register)

	return e.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}
