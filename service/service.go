package service

import (
	"github.com/google/uuid"
	"log"
	"message-service/model"
	"message-service/repository"
	"message-service/util"
	_ "message-service/util"
)

type MessageService interface {
	Create(message model.Message) model.Message
	GetAllMessages() []model.Message
	GetMessageById(messageId string) (model.Message, error)
	UpdateMessage(id string, msgText string) (model.Message, error)
	DeleteMessage(id string) error
}

type messageService struct {
	msgRepo repository.MessageRepository
}

func NewMessageService(msgRepo repository.MessageRepository) MessageService {
	return &messageService{
		msgRepo: msgRepo,
	}
}

func (s *messageService) Create(message model.Message) model.Message {
	log.Print(message.Text)
	msg := model.Message{
		Id:           uuid.New().String(),
		Text:         message.Text,
		IsPalindrome: util.IsPalindrome(message.Text),
	}
	return s.msgRepo.CreateMessage(msg)
}
func (s *messageService) GetAllMessages() []model.Message {
	return s.msgRepo.ListAllMessages()
}
func (s *messageService) GetMessageById(messageId string) (model.Message, error) {
	return s.msgRepo.GetMessageByID(messageId)
}

func (s *messageService) UpdateMessage(id string, messageText string) (model.Message, error) {
	msg := model.Message{
		Id:           id,
		Text:         messageText,
		IsPalindrome: util.IsPalindrome(messageText),
	}
	return s.msgRepo.UpdateMessage(id, msg)
}
func (s *messageService) DeleteMessage(id string) error {
	return s.msgRepo.DeleteMessage(id)
}
