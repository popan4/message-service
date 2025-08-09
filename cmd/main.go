package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"message-service/logger"

	"message-service/config"
	"message-service/controller"
	"message-service/repository"
	"message-service/router"
	"message-service/service"
)

func main() {

	// Initialize configuration
	loadConfig := config.LoadConfig()
	// Initialize Fiber app
	app := fiber.New()
	logger.Init()
	defer logger.Sync()

	logger.Info("Starting message service")
	messageRepository := repository.NewMessageRepository()
	messageService := service.NewMessageService(messageRepository)
	ctrl := controller.NewMessageController(messageService)

	// Setup routes
	router.SetupRoutes(app, ctrl)

	err := app.Listen(fmt.Sprintf(":%d", loadConfig.Server.Port))
	if err != nil {
		return
	}
}
