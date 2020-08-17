package main

import (
	"github.com/Angry-Lab/api/internal/env"
	"github.com/Angry-Lab/api/internal/server"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	config := &env.Config{}

	err := envconfig.Process("auth", config)
	if err != nil {
		logrus.Fatalf("process environment for geo users failed: %s", err)
	}

	logrus.Fatal(server.Run(config))
}
