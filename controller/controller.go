package controller

import (
	"github.com/gofiber/fiber/v2"
	"message-service/model"
	"message-service/service"
	"message-service/util"
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
	// Validate the message text
	if err := util.ValidateMessage(msgBody.Text); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	message := c.msgService.Create(msgBody)
	return ctx.Status(fiber.StatusCreated).JSON(message)
}

func (c *MessageController) GetMessages(ctx *fiber.Ctx) error {
	return ctx.JSON(c.msgService.GetAllMessages())
}

func (c *MessageController) GetMessageById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	message, err := c.msgService.GetMessageById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(message)
}

func (c *MessageController) UpdateMessage(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var body struct {
		Text string `json:"messageText"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	message, err := c.msgService.UpdateMessage(id, body.Text)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(message)
}

func (c *MessageController) DeleteMessage(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.msgService.DeleteMessage(id); err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
