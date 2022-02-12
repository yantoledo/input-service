package process_customer

import (
	"errors"

	"github.com/yantoledo/input-service/entity/customer"
	service "github.com/yantoledo/input-service/infra/service/customer_service"
)

type ProcessCustomer struct {
	Service service.CustomerServiceInterface
}

func NewProcessCustomer(service service.CustomerServiceInterface) *ProcessCustomer {
	return &ProcessCustomer{Service: service}
}

func (p *ProcessCustomer) Execute(input CustomerDtoInput) (CustomerDtoOutput, error) {

	customer := customer.NewCustomer()
	customer.Name = input.Name
	customer.UniqueID = input.UniqueID
	customer.UniqueClientID = input.UniqueClientID
	customer.Source = input.Source

	invalidCustomer := customer.IsValid()

	if invalidCustomer != nil {
		return CustomerDtoOutput{}, errors.New("Customer invalid")
	}

	idCustomer, err := p.Service.Insert(customer)
	if err != nil || idCustomer == 0 {
		return CustomerDtoOutput{}, errors.New("Customer's insert error")
	}

	output := CustomerDtoOutput{
		IdCustomer:     idCustomer,
		UniqueClientID: customer.UniqueClientID,
	}

	return output, nil
}
