package message_service

import "github.com/yantoledo/input-service/api/entity/message"

type MessageServiceMock struct {
}

func NewMessageServiceMock() *MessageServiceMock {
	return &MessageServiceMock{}
}

func (c *MessageServiceMock) Publish(message *message.Message) error {
	return nil
}
