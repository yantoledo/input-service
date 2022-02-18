package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yantoledo/input-service/api/service/customer_service"
	service "github.com/yantoledo/input-service/api/service/message_service"
)

func TestProcessMessageWhenItIsValid(t *testing.T) {
	customer := customer_service.CustomerProcessed{
		IdCustomer:     1,
		Name:           "John Lock",
		UniqueClientID: 1234,
		Source:         1,
	}

	input := MessageDtoInput{
		Text:     "Hi",
		Type:     "Text",
		MediaUrl: "https://www.url.com",
		Customer: customer,
	}

	expectedOutput := MessageDtoOutput{
		Text:     "Hi",
		Type:     "Text",
		MediaUrl: "https://www.url.com",
		Customer: customer,
	}

	serviceMock := service.NewMessageServiceMock()

	usecase := NewProcessMessage(serviceMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
