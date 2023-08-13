package config

import (
	"go-template/internal/pkg/database"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"go-template/internal/pkg/httpserver"
)

const (
	Local = "local"
	Prod  = "prod"
	Test  = "test"
)

type Config struct {
	Env      string                  `yaml:"env" env:"ENV" env-default:"prod"`
	Http     httpserver.HttpConfig   `yaml:"http_server"`
	DataBase database.DataBaseConfig `yaml:"database"`
}

func (c *Config) Parse() error {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "./config/local.yml"
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("config file does not exist: %s\n", configPath)

		return err
	}

	if err := cleanenv.ReadConfig(configPath, c); err != nil {
		return err
	}

	return nil
}
