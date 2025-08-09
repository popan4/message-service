package service

import (
	"log"
	"message-service/model"
)

type MessageService interface {
	Create(message model.Message) string
}

type messageService struct{}

func NewMessageService() MessageService {
	return &messageService{}
}

func (s *messageService) Create(message model.Message) string {
	log.Print(message.Text)
	//TODO validate the message
	// TODO save the message to the database
	// TODO check if the message is a palindrome
	return message.Text
}
