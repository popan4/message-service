package main

import (
  "github.com/gofiber/fiber/v2"
  "log"
  "message-service/controller"
  "message-service/router"
)

func main() {
  app := fiber.New()

  // Start the server on port 8080
  log.Println("Server is running on port 8080")
  ctrl := controller.NewMessageController()

  router.SetupRoutes(app, ctrl)
  err := app.Listen(":8080")
  if err != nil {
	return
  }

}
