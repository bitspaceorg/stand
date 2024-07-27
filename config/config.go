package config

import (
    "errors"
    "os"

    "github.com/joho/godotenv"
)

type AppConfig struct {
    ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
    if os.Getenv("APP_ENV") == "dev" {
        godotenv.Load()
    }

    httpPort := os.Getenv("HTTP_PORT")

    if len(httpPort) < 1 {
        return AppConfig{}, errors.New("Env variables not found")
    }

    return AppConfig{
        ServerPort: httpPort,
    }, nil
}
