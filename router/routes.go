package router

import (
	"github.com/gofiber/fiber/v2"
	"message-service/controller"
)

func SetupRoutes(app *fiber.App, messageController *controller.MessageController) {
	app.Get("/health", messageController.GetHealth)
	app.Get("/", messageController.GetHealth)
	app.Post("/createMsg", messageController.CreateMessage)
	app.Get("/getAllMsg", messageController.GetMessages)
	app.Get("/getMsg/:id", messageController.GetMessageById)
	app.Put("/updateMsg/:id", messageController.UpdateMessage)
	app.Delete("/deleteMsg/:id", messageController.DeleteMessage)
}
