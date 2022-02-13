package message

import (
	"errors"

	"github.com/yantoledo/input-service/entity/customer"
)

type Message struct {
	Text     string
	Type     string
	MediaUrl string
	Customer customer.Customer
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
