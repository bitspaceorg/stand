package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/deploy"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/puller"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/runtime"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
	"github.com/gofiber/fiber/v2"
)

type buildHandler struct {
}

type MessageStruct struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type BuildDTO struct {
	Link string `json:"link"`
	Name string `json:"name"`
}

func (b *buildHandler) build(c *fiber.Ctx) error {
	var dto BuildDTO
	if err := c.BodyParser(&dto); err != nil {
		return err
	}

	runtime := runtime.NodeRuntimeInstaller{}

	deployer := deploy.GetInstance(&runtime)

	projectFolder := fmt.Sprintf("%s/%s/", utils.ShadowFolder, dto.Name)
	buildFile := projectFolder + dto.Name + ".yml"

	if err := os.RemoveAll(projectFolder+dto.Name); err != nil {
		return err
	}

	err := deployer.Kill(dto.Name)
	if err != nil {
		return err
	}

	puller := puller.GitPuller{
		RepoLink: dto.Link,
		Path:     projectFolder + dto.Name,
	}

	err = puller.Pull()
	if err != nil {
		return err
	}

	msgChan := make(chan MessageStruct)
	var msg MessageStruct
	go deployer.Deploy(buildFile, func(message string, success bool) {
		msgChan <- MessageStruct{
			Message: message,
			Success: success,
		}
		close(msgChan)
	})
	msg = <-msgChan
	log.Println(msg)
	return c.JSON(msg)
}

func SetupBuildRoutes(rh *rest.RestHandler) {
	app := rh.App
	b := buildHandler{}
	app.Post("/build", b.build)
}
