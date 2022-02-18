package usecase

import (
	"errors"

	"github.com/yantoledo/input-service/api/entity/message"
	service "github.com/yantoledo/input-service/api/service/message_service"
)

type ProcessMessage struct {
	Service service.MessageServiceInterface
}

func NewProcessMessage(service service.MessageServiceInterface) *ProcessMessage {
	return &ProcessMessage{Service: service}
}

func (p *ProcessMessage) Execute(input MessageDtoInput) (MessageDtoOutput, error) {

	message := message.NewMessage()
	message.Text = input.Text
	message.Type = input.Type
	message.MediaUrl = input.MediaUrl
	message.Customer = input.Customer

	invalidMessage := message.IsValid()

	if invalidMessage != nil {
		return MessageDtoOutput{}, errors.New("Invalid message")
	}

	err := p.Service.Publish(message)
	if err != nil {
		return MessageDtoOutput{}, errors.New("Publish message error")
	}

	output := MessageDtoOutput{
		Text:     message.Text,
		Type:     message.Type,
		MediaUrl: message.MediaUrl,
		Customer: message.Customer,
	}

	return output, nil
}
