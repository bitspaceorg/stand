package config

import (
	"os"

	"github.com/joho/godotenv"
)

const httpPort = "6789"

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	return AppConfig{
		ServerPort: httpPort,
	}, nil
}
