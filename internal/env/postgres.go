package env

import "time"

type Postgres struct {
	DSN          string        `envconfig:"postgres_dsn" default:""`
	ConnLifetime time.Duration `envconfig:"postgres_conn_lifetime" default:"2s"`
}
