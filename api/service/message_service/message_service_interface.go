package service

import "github.com/yantoledo/input-service/api/entity/message"

type MessageServiceInterface interface {
	Publish(message *message.Message) error
}
