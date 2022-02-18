package message

import (
	"errors"

	"github.com/yantoledo/input-service/api/service/customer_service"
)

type Message struct {
	Text     string
	Type     string
	MediaUrl string
	Customer customer_service.CustomerProcessed
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) IsValid() error {
	if m.Text == "" || m.Type == "" {
		return errors.New("Invalid Message")
	}

	return nil
}
