package controller

import "github.com/gofiber/fiber/v2"

type MessageController struct {
}

func NewMessageController() *MessageController {
	return &MessageController{}
}

func (c *MessageController) GetHealth(ctx *fiber.Ctx) error {
	return ctx.SendString("SERVICE IS UP AND RUNNING")
}
