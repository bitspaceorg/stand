package handlers

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
)

type repoHandler struct {
}

func (h *repoHandler) getRepo(c *fiber.Ctx) error {
	file, err := os.Open(utils.GetShadowReposFilePath())
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"repos": arr,
	})
}

func (h *repoHandler) repo(c *fiber.Ctx) error {
	body := struct {
		Repo []string `json:"repos"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	for _, repo := range body.Repo {
		err := utils.AppendToFile(utils.GetShadowReposFilePath(), fmt.Sprintf("%v\n", repo))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func SetupRepoRoutes(rh *rest.RestHandler) {
	app := rh.App
	h := repoHandler{}
	app.Post("/repos", h.repo)
	app.Get("/repos", h.getRepo)
}
