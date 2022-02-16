package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yantoledo/input-service/entity/customer"
	service "github.com/yantoledo/input-service/infra/service/message_service"
)

func TestProcessMessageWhenItIsValid(t *testing.T) {
	customer := customer.NewCustomer()
	customer.Name = "John Lock"
	customer.UniqueID = 1234
	customer.UniqueClientID = 1234
	customer.Source = 1

	input := MessageDtoInput{
		Text:     "Hi",
		Type:     "Text",
		MediaUrl: "https://www.url.com",
		Customer: *customer,
	}

	expectedOutput := MessageDtoOutput{
		Text:     "Hi",
		Type:     "Text",
		MediaUrl: "https://www.url.com",
		Customer: *customer,
	}

	serviceMock := service.NewMessageServiceMock()

	usecase := NewProcessMessage(serviceMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
