package message_service

import (
	"encoding/json"

	"github.com/yantoledo/input-service/api/entity/message"
	"github.com/yantoledo/input-service/infra/message_broker"
)

type MessageService struct {
	BrokerClient message_broker.BrokerInterface
}

func NewMessageService(broker message_broker.BrokerInterface) *MessageService {
	return &MessageService{BrokerClient: broker}
}

func (m *MessageService) Publish(message *message.Message) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	brokerRequest := message_broker.BrokerRequest{
		URL:  "https://7a3f-186-193-220-142.ngrok.io", // TODO: We must get the url from a env variable in config file
		Body: body,
	}

	err = m.BrokerClient.Publish(brokerRequest)

	if err != nil {
		return err
	}

	return nil
}
