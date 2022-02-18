package customer_service

import (
	"github.com/yantoledo/input-service/api/entity/customer"
)

type CustomerServiceInterface interface {
	Insert(customer *customer.Customer) (CustomerProcessed, error)
}
