package usecase

import (
	"errors"

	"github.com/yantoledo/input-service/api/entity/customer"
	service "github.com/yantoledo/input-service/api/service/customer_service"
)

type ProcessCustomer struct {
	Service service.CustomerServiceInterface
}

func NewProcessCustomer(service service.CustomerServiceInterface) *ProcessCustomer {
	return &ProcessCustomer{Service: service}
}

func (p *ProcessCustomer) Execute(input CustomerDtoInput) (CustomerDtoOutput, error) {

	customer := customer.NewCustomer()
	customer.Name = "Input"
	customer.UniqueID = input.UniqueID
	customer.UniqueClientID = input.UniqueClientID
	customer.Source = input.Source

	invalidCustomer := customer.IsValid()

	if invalidCustomer != nil {
		return CustomerDtoOutput{}, errors.New("Invalid Customer")
	}

	service_response, err := p.Service.Insert(customer)
	if err != nil || service_response.IdCustomer == 0 {
		return CustomerDtoOutput{}, errors.New("Customer insert error")
	}

	dtoOutput := CustomerDtoOutput{
		IdCustomer:     service_response.IdCustomer,
		Name:           service_response.Name,
		UniqueClientID: service_response.UniqueClientID,
	}

	return dtoOutput, nil
}
