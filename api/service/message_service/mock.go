package message_service

import (
	"github.com/yantoledo/input-service/api/entity/message"
	"github.com/yantoledo/input-service/infra/message_broker"
)

type MessageServiceMock struct {
	BrokerClient message_broker.BrokerInterface
}

func NewMessageServiceMock(brokerClient message_broker.BrokerInterface) *MessageServiceMock {
	return &MessageServiceMock{BrokerClient: brokerClient}
}

func (c *MessageServiceMock) Publish(message *message.Message) error {
	return nil
}
