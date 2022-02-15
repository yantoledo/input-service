package service

import (
	"github.com/yantoledo/input-service/entity/customer"
)

type CustomerServiceMock struct {
}

func NewCustomerServiceMock() *CustomerServiceMock {
	return &CustomerServiceMock{}
}

func (c *CustomerServiceMock) Insert(customer *customer.Customer) (int, error) {
	return 1, nil
}
