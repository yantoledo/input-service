package usecase

import (
	service "github.com/yantoledo/input-service/api/service/customer_service"
)

type MessageDtoInput struct {
	Text     string                    `json:"text"`
	Type     string                    `json:"type"`
	MediaUrl string                    `json:"midia_url"`
	Customer service.CustomerProcessed `json:"customer"`
}

type MessageDtoOutput struct {
	Text     string                    `json:"text"`
	Type     string                    `json:"type"`
	MediaUrl string                    `json:"midia_url"`
	Customer service.CustomerProcessed `json:"customer"`
}
