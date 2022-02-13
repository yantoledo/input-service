package service

import (
	"net/http"

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

func (c *CustomerServiceMock) GetHeader() http.Header {
	key := "Content_type"
	value := []string{"application/json"}
	header := http.Header(map[string][]string{key: value})

	return header
}
