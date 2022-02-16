package usecase

import "github.com/yantoledo/input-service/entity/customer"

type MessageDtoInput struct {
	Text     string            `json:"text"`
	Type     string            `json:"type"`
	MediaUrl string            `json:"midia_url"`
	Customer customer.Customer `json:"customer"`
}

type MessageDtoOutput struct {
	Text     string            `json:"text"`
	Type     string            `json:"type"`
	MediaUrl string            `json:"midia_url"`
	Customer customer.Customer `json:"customer"`
}
