package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
)

type healthHandler struct {
}

func (h *healthHandler) health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func SetupHealthRoutes(rh *rest.RestHandler) {
	app := rh.App
	h := healthHandler{}
	app.Get("/health", h.health)
}
