package message

import (
	"errors"

	"github.com/yantoledo/input-service/api/usecase/process_customer"
)

type Message struct {
	Text     string
	Type     string
	MediaUrl string
	Customer process_customer.CustomerDtoOutput
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
