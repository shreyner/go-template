package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port  int    `env:"PORT" envDefault:"8080"`
	DBUrl string `env:"DATABASE_URL,required"`
}

func (c *Config) Parse() error {
	if err := env.Parse(c); err != nil {
		return err
	}

	return nil
}
