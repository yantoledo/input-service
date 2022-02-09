package customer

import "errors"

type Customer struct {
	Name           string
	UniqueID       int
	UniqueClientID int
	Source         int
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c *Customer) IsValid() error {
	if c.Name == "" {
		return errors.New("Customer's name is empty")
	}

	if c.UniqueID == 0 || c.UniqueClientID == 0 || c.Source == 0 {
		return errors.New("Customer is invalid")
	}

	return nil
}
