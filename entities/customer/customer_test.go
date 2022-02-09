package customer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerWithoutName(t *testing.T) {
	//Without Name
	customer := NewCustomer()
	customer.Name = ""
	customer.UniqueID = 123456
	customer.UniqueClientID = 122
	customer.Source = 1

	err := customer.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Customer's name is empty", err.Error())
}

func TestCustomerWithoutUniqueID(t *testing.T) {
	customer := NewCustomer()
	customer.Name = "John Lock"
	customer.UniqueID = 0
	customer.UniqueClientID = 122
	customer.Source = 1

	err := customer.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Customer is invalid", err.Error())
}

func TestCustomerWithoutUniqueClientID(t *testing.T) {
	customer := NewCustomer()
	customer.Name = "John Lock"
	customer.UniqueID = 123456
	customer.UniqueClientID = 0
	customer.Source = 1

	err := customer.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Customer is invalid", err.Error())
}

func TestCustomerWithoutSource(t *testing.T) {
	customer := NewCustomer()
	customer.Name = "John Lock"
	customer.UniqueID = 123456
	customer.UniqueClientID = 122
	customer.Source = 0

	err := customer.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Customer is invalid", err.Error())
}
