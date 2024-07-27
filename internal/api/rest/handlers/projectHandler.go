package handlers

import (
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	parser "github.com/bitspaceorg/STAND-FOSSHACK/internal/build-parser"
	"github.com/gofiber/fiber/v2"
)

type projectHandler struct {
}

func (h *projectHandler) newProject(c *fiber.Ctx) error {
	var cfg parser.NodeBuildConfig
	if err := c.BodyParser(cfg); err != nil {
	} 
	return nil
}

func SetupProjectRoutes(rh *rest.RestHandler) {
	// app := rh.App
	// h := projectHandler{}
}
