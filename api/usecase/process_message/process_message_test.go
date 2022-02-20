package process_message

import (
	"testing"

	"github.com/stretchr/testify/assert"
	service "github.com/yantoledo/input-service/api/service/message_service"
	"github.com/yantoledo/input-service/api/usecase/process_customer"
	"github.com/yantoledo/input-service/infra/message_broker"
)

func TestProcessMessageWhenItIsValid(t *testing.T) {
	customer := process_customer.CustomerDtoOutput{
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

	brokerMock := message_broker.NewBrokerClientMock()

	serviceMock := service.NewMessageServiceMock(brokerMock)

	usecase := NewProcessMessage(serviceMock)
	err := usecase.Execute(input)

	assert.Nil(t, err)

}
