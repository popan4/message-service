package repository

import (
	"errors"
	"message-service/model"
)

type MessageRepository interface {
	CreateMessage(message model.Message) model.Message
	ListAllMessages() []model.Message
	GetMessageByID(messageId string) (model.Message, error)
	UpdateMessage(id string, msg model.Message) (model.Message, error)
	DeleteMessage(id string) error
}

type messageRepository struct {
	messagesdb map[string]model.Message
}

func (mr messageRepository) GetMessageByID(msgId string) (model.Message, error) {
	if msg, exists := mr.messagesdb[msgId]; exists {
		return msg, nil
	}
	return model.Message{}, errors.New("message not found")
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{
		messagesdb: make(map[string]model.Message),
	}
}

func (mr messageRepository) CreateMessage(message model.Message) model.Message {
	mr.messagesdb[message.Id] = message
	return message
}

func (mr messageRepository) ListAllMessages() []model.Message {
	listMsgs := []model.Message{}
	for _, msg := range mr.messagesdb {
		listMsgs = append(listMsgs, msg)
	}
	return listMsgs
}

func (mr messageRepository) UpdateMessage(id string, msg model.Message) (model.Message, error) {
	if _, exists := mr.messagesdb[id]; !exists {
		return model.Message{}, errors.New("message not found")
	}
	msg.Id = id
	mr.messagesdb[id] = msg
	return msg, nil
}

func (mr messageRepository) DeleteMessage(id string) error {
	if _, exists := mr.messagesdb[id]; !exists {
		return errors.New("message not found")
	}
	delete(mr.messagesdb, id)
	return nil
}
