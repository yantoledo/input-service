package process_message

import (
	"github.com/yantoledo/input-service/api/usecase/process_customer"
)

type MessageDtoInput struct {
	Text     string                             `json:"text"`
	Type     string                             `json:"type"`
	MediaUrl string                             `json:"midia_url"`
	Customer process_customer.CustomerDtoOutput `json:"customer"`
}

type MessageDtoOutput struct {
	Text     string                             `json:"text"`
	Type     string                             `json:"type"`
	MediaUrl string                             `json:"midia_url"`
	Customer process_customer.CustomerDtoOutput `json:"customer"`
}
