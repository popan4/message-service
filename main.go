package main

import (
  "github.com/gofiber/fiber/v2"
  "log"
  "message-service/controller"
  "message-service/router"
  "message-service/service"
)

func main() {
  app := fiber.New()

  // Start the server on port 8080
  log.Println("Server is running on port 8080")

  messageService := service.NewMessageService()
  ctrl := controller.NewMessageController(messageService)

  // Setup routes
  router.SetupRoutes(app, ctrl)
  err := app.Listen(":8080")
  if err != nil {
	return
  }

}
