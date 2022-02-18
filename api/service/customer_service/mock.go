package customer_service

import (
	"github.com/yantoledo/input-service/api/entity/customer"
)

type CustomerServiceMock struct {
}

func NewCustomerServiceMock() *CustomerServiceMock {
	return &CustomerServiceMock{}
}

func (c *CustomerServiceMock) Insert(customer *customer.Customer) (CustomerProcessed, error) {
	return CustomerProcessed{IdCustomer: 1, Name: "John Lock", UniqueClientID: 12345, Source: 1}, nil
}
