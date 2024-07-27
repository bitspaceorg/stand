package handlers

import (
	"log"
	"net/http"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/bitspaceorg/STAND-FOSSHACK/internal/deploy"
	"github.com/gofiber/fiber/v2"
)

type buildHandler struct {
}

type MessageStruct struct {
	Message string `json:"message"`
	Success bool `json:"success"`
}
func (b *buildHandler) build(c *fiber.Ctx) error {
	msgChan := make(chan MessageStruct)
	var msg MessageStruct
	go deploy.DeployGo("/home/suryaprakash/Dev/Go/STAND/example/test.yml", func(message string, success bool) {
		msgChan <- MessageStruct{
			Message: message,
			Success: success,
		}
		close(msgChan)
	})
	msg = <- msgChan
	log.Println(msg)
	return c.JSON(msg)
}

// func SetupBuildRoutes(rh *rest.RestHandler) {
// 	app := rh.App
// 	b := buildHandler{}
// 	app.Get("/build", b.build)
// =======
//
// type buildHandler struct {
// }
//
//
// func (b *buildHandler) build(c *fiber.Ctx) error {
//     deploy.DeployGo("/home/t-aswath/projects/STAND-FOSSHACK/example/test.yml",func(message string,success bool){
//         if success {
//             c.Status(http.StatusOK).JSON(fiber.Map{
//                 "status": "ok",
//                 "message": message,
//             })
//         } else {
//             c.Status(http.StatusInternalServerError).JSON(fiber.Map{
//                 "status": "error",
//                 "message": message,
//             })
//         }
//     
//     });
//     return c.Status(http.StatusOK).JSON(fiber.Map{
//         "status": "ok",
//     })
// }

func SetupBuildRoutes(rh *rest.RestHandler) {
    app := rh.App
    b := buildHandler{}
    app.Get("/build", b.build)
}
