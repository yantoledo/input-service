package process_customer

import (
	"errors"

	"github.com/yantoledo/input-service/entity"
	"github.com/yantoledo/input-service/entity/customer"
)

type ProcessCustomer struct {
	Repository entity.CustomerRepository
}

func NewProcessCustomer(repository entity.CustomerRepository) *ProcessCustomer {
	return &ProcessCustomer{Repository: repository}
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

	err := p.Repository.Insert(customer.Name, customer.UniqueID, customer.UniqueClientID, customer.Source)
	if err != nil {
		return CustomerDtoOutput{}, errors.New("Customer's insert error")
	}

	output := CustomerDtoOutput{
		UniqueID:       customer.UniqueID,
		UniqueClientID: customer.UniqueClientID,
	}

	return output, nil
}
