package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/Angry-Lab/api/internal/env"
	"github.com/Angry-Lab/api/pkg/user"
	"github.com/Angry-Lab/api/pkg/user/postgres"
	"github.com/spf13/pflag"
)

var (
	login    = pflag.String("login", "", "user login")
	name     = pflag.String("name", "", "user name")
	password = pflag.String("password", "", "user password")
	role     = pflag.String("role", "", "user role")
	prefix   = pflag.String("prefix ", "cli", "environment prefix")
)

func main() {
	pflag.Parse()

	config, err := env.ReadConfig(*prefix)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", config.Postgres.DSN)
	if err != nil {
		log.Fatalf("open sql connection failed: %s", err)
	}

	db.SetConnMaxLifetime(config.Postgres.ConnLifetime)

	u := user.User{
		Login:     *login,
		Password:  *password,
		Status:    user.StatusActive,
		Name:      *name,
		Role:      user.Role(*role),
		DTCreated: time.Now().UTC(),
		DTUpdated: time.Now().UTC(),
	}

	userRepo := postgres.Users(db)
	tokenRepo := postgres.Tokens(db)
	users := user.NewUseCase(userRepo, tokenRepo)

	err = users.Register(context.Background(), &u)
	if err != nil {
		log.Fatal(err)
	}
}
