package service

import (
	"github.com/yantoledo/input-service/entity/message"
	"github.com/yantoledo/input-service/infra/message_broker"
)

type MessageService struct {
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (m *MessageService) Publish(message *message.Message) error {
	url := "https://7a3f-186-193-220-142.ngrok.io" // TODO: We must get the url from a env variable in config file

	messageBroker := message_broker.NewPublisher()
	err := messageBroker.Publish(url, message)

	if err != nil {
		return err
	}

	return nil
}
