package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/bitspaceorg/STAND-FOSSHACK/utils"
	"github.com/gofiber/fiber/v2"
)

type userVerification struct {
}

func ValidateUser(username string, password string) bool {
	hashedPassword := utils.HashPassword(password)

	file, err := os.Open(utils.GetShadowAuthFilePath())
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid shadow file format.")
			return false
		}

		if parts[0] == username && parts[1] == hashedPassword {
			fmt.Println("User authenticated.")
			return true
		}
	}

	fmt.Println("Invalid username or password.")
	return false
}

func (h *userVerification) verifyUser(c *fiber.Ctx) error {
	username := c.Query("username")
	password := c.Query("password")
	status := false
	if ValidateUser(username, password) {
		status = true
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": status,
	})
}

func SetupUserVerificationRoutes(rh *rest.RestHandler) {
	app := rh.App
	h := userVerification{}
	app.Post("/verifyUser", h.verifyUser)
}
