package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Angry-Lab/api/internal/env"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/Angry-Lab/api/pkg/user/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func Run(config *env.Config) error {
	db, err := sql.Open("postgres", config.Postgres.DSN)
	if err != nil {
		return errors.Errorf("open sql connection failed: %s", err)
	}

	db.SetConnMaxLifetime(config.Postgres.ConnLifetime)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodPost, http.MethodGet, http.MethodOptions},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept"},
		AllowCredentials: true,
	}))

	userRepo := postgres.Users(db)
	tokenRepo := postgres.Tokens(db)

	h := &Handler{
		Users: user.NewUseCase(userRepo, tokenRepo),
	}

	e.POST("/v1/auth", h.Auth)
	e.POST("/v1/auth/reset", h.ResetPassword)

	e.GET("/v1/users", h.ListUsers, h.WithAuth, h.WithAccess(user.RoleOwner))
	e.POST("/v1/users", h.CreateUser, h.WithAuth, h.WithAccess(user.RoleOwner))
	e.PUT("/v1/users", h.UpdateUser, h.WithAuth, h.WithAccess(user.RoleOwner))
	e.DELETE("/v1/users", h.DeleteUser, h.WithAuth, h.WithAccess(user.RoleOwner))

	return e.Start(fmt.Sprintf("%s:%d", config.Host, config.Port))
}
