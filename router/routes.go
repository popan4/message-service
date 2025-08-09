package router

import (
	"github.com/gofiber/fiber/v2"
	"message-service/controller"
)

func SetupRoutes(app *fiber.App, messageController *controller.MessageController) {
	app.Get("/health", messageController.GetHealth)
}
