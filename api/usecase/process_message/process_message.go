package process_message

import (
	"github.com/yantoledo/input-service/api/entity/message"
	service "github.com/yantoledo/input-service/api/service/message_service"
)

type ProcessMessage struct {
	Service service.MessageServiceInterface
}

func NewProcessMessage(service service.MessageServiceInterface) *ProcessMessage {
	return &ProcessMessage{Service: service}
}

func (p *ProcessMessage) Execute(input MessageDtoInput) error {

	message := message.NewMessage()
	message.Text = input.Text
	message.Type = input.Type
	message.MediaUrl = input.MediaUrl
	message.Customer = input.Customer

	err := message.IsValid()

	if err != nil {
		return err
	}

	err = p.Service.Publish(message)
	if err != nil {
		return err
	}

	return nil
}
