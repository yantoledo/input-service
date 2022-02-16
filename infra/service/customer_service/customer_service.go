package service

import (
	"encoding/json"
	"fmt"

	"github.com/yantoledo/input-service/entity/customer"
	"github.com/yantoledo/input-service/infra/http_protocol"
)

type CustomerService struct {
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (c *CustomerService) Insert(customer *customer.Customer) (CustomerServiceOutput, error) {
	body, err := json.Marshal(customer)

	if err != nil {
		return CustomerServiceOutput{}, err
	}

	httpService := http_protocol.NewHttpService()

	res, err := httpService.Post(http_protocol.HTTPRequest{
		URL:  "https://69fc-186-193-220-142.ngrok.io", // TODO: We must get the url from an env variable in config file
		Body: body,
	})
	fmt.Println(res) //TODO: Removes after
	if err != nil {
		return CustomerServiceOutput{}, err
	}

	output := CustomerServiceOutput{
		IdCustomer:     1, //TODO: We must return idCustomer from response body
		Name:           "John Lock",
		UniqueClientID: customer.UniqueClientID,
	}

	return output, nil
}
