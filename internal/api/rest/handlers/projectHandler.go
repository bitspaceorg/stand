package handlers

import (
	"os"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	parser "github.com/bitspaceorg/STAND-FOSSHACK/internal/build-parser"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type projectHandler struct {
}

func (h *projectHandler) newProject(c *fiber.Ctx) error {
	var cfg parser.NodeBuildConfig
	if err := c.BodyParser(&cfg); err != nil {
		return err
	}
	ymlData, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	err = os.Mkdir(cfg.Project.Home+"/", 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(cfg.Project.Home + "/" + cfg.Project.Name + ".yml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(ymlData)

	if err != nil {
		return err
	}

	return nil
}

func SetupProjectRoutes(rh *rest.RestHandler) {
	app := rh.App
	h := projectHandler{}

	app.Post("/newProject", h.newProject)
}
