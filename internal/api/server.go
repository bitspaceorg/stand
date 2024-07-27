package api

import (
	"github.com/bitspaceorg/STAND-FOSSHACK/config"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest/handlers"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})
	app.Use(cors.New())

	rh := rest.RestHandler{
		App: app,
	}

	app.Static("/", utils.GetShadowGUIBuildPath())

	setupRoutes(&rh)

	err := app.Listen(":" + config.ServerPort)

	if err != nil {
		log.Fatalf("Error starting the server :%v", err)
	}
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupHealthRoutes(rh)
	handlers.SetupMetricRoutes(rh)
	// handlers.SetupUserRoutes(rh)
}
