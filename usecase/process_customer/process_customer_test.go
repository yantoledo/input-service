package process_customer

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	service "github.com/yantoledo/input-service/infra/service/customer_service"
)

func TestProcessCustomerWhenItIsValid(t *testing.T) {
	input := CustomerDtoInput{
		Name:           "John Lock",
		UniqueID:       1234,
		UniqueClientID: 1234,
		Source:         1,
	}

	expectedOutput := CustomerDtoOutput{
		idCustomer:     1,
		UniqueClientID: 1234,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := service.NewCustomerServiceMock()

	usecase := NewProcessCustomer(serviceMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
