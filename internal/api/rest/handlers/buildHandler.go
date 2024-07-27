package handlers

import (
	"net/http"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/deploy"
	"github.com/gofiber/fiber/v2"
)


type buildHandler struct {
}


func (b *buildHandler) build(c *fiber.Ctx) error {
    deploy.DeployGo("/home/t-aswath/projects/STAND-FOSSHACK/example/test.yml");
    return c.Status(http.StatusOK).JSON(fiber.Map{
        "status": "ok",
    })
}

func SetupBuildRoutes(rh *rest.RestHandler) {
    app := rh.App
    b := buildHandler{}
    app.Get("/build", b.build)
}
