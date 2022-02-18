package message_service

import (
	"github.com/yantoledo/input-service/api/entity/message"
	"github.com/yantoledo/input-service/infra/message_broker"
)

type MessageService struct {
	PBClient message_broker.PublisherInterface
}

func NewMessageService() *MessageService {
	return &MessageService{}
}

func (m *MessageService) Publish(message *message.Message) error {
	url := "https://7a3f-186-193-220-142.ngrok.io" // TODO: We must get the url from a env variable in config file

	err := m.PBClient.Publish(url, message)

	if err != nil {
		return err
	}

	return nil
}
