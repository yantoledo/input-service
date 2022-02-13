package service

import "github.com/yantoledo/input-service/entity/message"

type MessageServiceInterface interface {
	Publish(message *message.Message) error
}
