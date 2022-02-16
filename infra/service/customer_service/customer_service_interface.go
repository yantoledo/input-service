package service

import (
	"github.com/yantoledo/input-service/entity/customer"
)

type CustomerServiceInterface interface {
	Insert(customer *customer.Customer) (CustomerServiceOutput, error)
}
