package service

import (
	"net/http"

	"github.com/yantoledo/input-service/entity/customer"
)

type CustomerServiceInterface interface {
	GetHeader() http.Header
	Insert(customer *customer.Customer) (int, error)
}
