package main

import (
	_ "github.com/lib/pq"

	"github.com/Angry-Lab/api/internal/env"
	"github.com/Angry-Lab/api/internal/server"
	"github.com/labstack/gommon/log"
	"github.com/spf13/pflag"
)

var (
	prefix = pflag.String("prefix ", "api", "environment prefix")
)

func main() {
	pflag.Parse()

	logger := log.New(*prefix)

	config, err := env.ReadConfig(*prefix)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Fatal(server.Run(config))
}
