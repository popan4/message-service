package router

import (
  "github.com/gofiber/fiber/v2"
  "message-service/controller"
)

func SetupRoutes(app *fiber.App, messageController *controller.MessageController) {
  app.Get("/health", messageController.GetHealth)
  app.Get("/", messageController.GetHealth)
  app.Post("/create", messageController.CreateMessage)
  //TODO GET, LOAD, UPDATE, DELETE , list all messages

}
