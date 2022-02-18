package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yantoledo/input-service/entity/customer"
	"github.com/yantoledo/input-service/infra/http_protocol"
)

func TestCustomerService(t *testing.T) {

	customer := customer.NewCustomer()
	customer.Name = ""
	customer.UniqueID = 123456
	customer.UniqueClientID = 122
	customer.Source = 1

	expectedOut := CustomerServiceOutput{
		IdCustomer:     1,
		Name:           "John Lock",
		UniqueClientID: customer.UniqueClientID,
	}

	httpMock := http_protocol.NewHttpClientMock()

	service := NewCustomerService(httpMock)

	output, err := service.Insert(customer)
	assert.Nil(t, err)
	assert.Equal(t, expectedOut, output)

}
