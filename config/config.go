package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	ErrOnLoad    = errors.New("failed to load env from file")
	ErrOnProcess = errors.New("failed to process env")
)

type Config struct {
	LogLevel   string `envconfig:"LOG_LEVEL" default:"debug"`
	PgDSN      string `envconfig:"PG_DSN" required:"true"`
	HTTPListen string `envconfig:"HTTP_LISTEN" default:":8000"`
}

func Load(filenames ...string) (Config, error) {
	// by default loads from .env
	err := godotenv.Load(filenames...)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Config{}, errors.Join(ErrOnLoad, err)
	}

	conf := Config{}
	err = envconfig.Process("", &conf)
	if err != nil {
		return Config{}, errors.Join(ErrOnProcess, err)
	}

	return conf, nil
}
