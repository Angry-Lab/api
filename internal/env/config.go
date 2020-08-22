package env

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Host string `envconfig:"host" default:"0.0.0.0" json:"host"`
	Port uint16 `envconfig:"port" default:"6363" json:"port"`

	Postgres Postgres
}

func ReadConfig(prefix string) (*Config, error) {
	config := Config{}

	err := envconfig.Process(prefix, &config)
	if err != nil {
		return nil, errors.Wrap(err, "read config from environment failed")
	}

	err = envconfig.Process(prefix, &config.Postgres)
	if err != nil {
		return nil, errors.Wrap(err, "read postgres config from environment failed")
	}

	return &config, nil
}
