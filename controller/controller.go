package controller

import (
	"github.com/gofiber/fiber/v2"
	"message-service/model"
	"message-service/service"
)

type MessageController struct {
	msgService service.MessageService
}

func NewMessageController(messageService service.MessageService) *MessageController {
	return &MessageController{
		msgService: messageService,
	}
}

func (c *MessageController) GetHealth(ctx *fiber.Ctx) error {
	return ctx.SendString("SERVICE IS UP AND RUNNING")
}

func (c *MessageController) CreateMessage(ctx *fiber.Ctx) error {
	var msgBody model.Message
	if err := ctx.BodyParser(&msgBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	message := c.msgService.Create(msgBody)
	return ctx.Status(fiber.StatusCreated).JSON(message)
}
