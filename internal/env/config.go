package env

type Config struct {
	Host string `envconfig:"host" default:"0.0.0.0" json:"host"`
	Port uint16 `envconfig:"port" default:"6363" json:"port"`
}
