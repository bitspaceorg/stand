package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	parser "github.com/bitspaceorg/STAND-FOSSHACK/internal/build-parser"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/deploy"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/puller"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/runtime"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type projectHandler struct {
}

func (h *projectHandler) newProject(c *fiber.Ctx) error {
	var cfg parser.NodeBuildConfig

	runtime := runtime.NodeRuntimeInstaller{}

	deployer :=  deploy.GetInstance(&runtime)


	if err := c.BodyParser(&cfg); err != nil {
		return err
	}
	ymlData, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	projectFolder := fmt.Sprintf("%s/%s/", utils.ShadowFolder, cfg.Project.Name)
	buildFile := projectFolder + cfg.Project.Name + ".yml"
	err = os.MkdirAll(projectFolder, 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(buildFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(ymlData)

	if err != nil {
		return err
	}

	puller := puller.GitPuller{
		RepoLink: cfg.Project.RepoLink,
		Path:     projectFolder + cfg.Project.Name,
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

func SetupProjectRoutes(rh *rest.RestHandler) {
	app := rh.App
	h := projectHandler{}

	app.Post("/newProject", h.newProject)
}
