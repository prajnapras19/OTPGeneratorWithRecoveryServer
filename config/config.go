package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RESTPort string `envconfig:"REST_PORT" default:"80"`

	MySQLConfig MySQLConfig
}

type MySQLConfig struct {
	Username string `envconfig:"MYSQL_USER" default:""`
	Password string `envconfig:"MYSQL_PASSWORD" default:""`
	Hostname string `envconfig:"MYSQL_HOST" default:""`
	DBName   string `envconfig:"MYSQL_DATABASE" default:""`

	// config below can have a default value
	Port      string `envconfig:"MYSQL_PORT" default:"3306"`
	Charset   string `envconfig:"MYSQL_CHARSET" default:"utf8mb4"`
	ParseTime string `envconfig:"MYSQL_PARSETIME" default:"true"`
	Loc       string `envconfig:"MYSQL_LOC" default:"Local"`
}

func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
