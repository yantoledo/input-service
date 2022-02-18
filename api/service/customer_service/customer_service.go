package service

import (
	"encoding/json"
	"fmt"

	"github.com/yantoledo/input-service/api/entity/customer"
	"github.com/yantoledo/input-service/infra/http_protocol"
)

type CustomerService struct {
	HttpClient http_protocol.HttpClientInterface
}

func NewCustomerService(httpClient http_protocol.HttpClientInterface) *CustomerService {
	return &CustomerService{HttpClient: httpClient}
}

func (c *CustomerService) Insert(customer *customer.Customer) (CustomerServiceOutput, error) {
	body, err := json.Marshal(customer)

	if err != nil {
		return CustomerServiceOutput{}, err
	}

	res, err := c.HttpClient.Post(http_protocol.HTTPRequest{
		URL:  "https://69fc-186-193-220-142.ngrok.io", // TODO: We must get the url from an env variable in config file
		Body: body,
	})
	fmt.Println(res) //TODO: Removes after
	if err != nil {
		return CustomerServiceOutput{}, err
	}

	output := CustomerServiceOutput{
		IdCustomer:     1,           //TODO: We must get idCustomer from response body
		Name:           "John Lock", //TODO: We must get Name from response body
		UniqueClientID: customer.UniqueClientID,
	}

	return output, nil
}
