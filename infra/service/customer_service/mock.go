package service

import (
	"github.com/yantoledo/input-service/entity/customer"
)

type CustomerServiceMock struct {
}

func NewCustomerServiceMock() *CustomerServiceMock {
	return &CustomerServiceMock{}
}

func (c *CustomerServiceMock) Insert(customer *customer.Customer) (CustomerServiceOutput, error) {
	return CustomerServiceOutput{IdCustomer: 1, Name: "John Lock", UniqueClientID: 12345}, nil
}
