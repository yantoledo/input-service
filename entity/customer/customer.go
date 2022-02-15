package customer

import "errors"

type Customer struct {
	Name           string `json:"name"`
	UniqueID       int    `json:"unique_id"`
	UniqueClientID int    `json:"unique_client_id"`
	Source         int    `json:"source"`
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
