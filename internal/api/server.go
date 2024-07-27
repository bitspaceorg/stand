package api

import (
    "github.com/bitspaceorg/STAND/config"
    "github.com/bitspaceorg/STAND/internal/api/rest"
    "github.com/bitspaceorg/STAND/internal/api/rest/handlers"
    "log"

    "github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
    app := fiber.New()

    rh := rest.RestHandler{
        App: app,
    }
    setupRoutes(&rh)

    err := app.Listen(":" + config.ServerPort)

    if err != nil {
        log.Fatalf("Error starting the server :%v", err)
    }
}

func setupRoutes(rh *rest.RestHandler) {
    handlers.SetupHealthRoutes(rh)
    // handlers.SetupUserRoutes(rh)
}
