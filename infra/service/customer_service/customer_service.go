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

func (c *CustomerService) Insert(customer *customer.Customer) (int, error) {
	body, err := json.Marshal(customer)

	if err != nil {
		return 0, err
	}

	httpService := http_protocol.NewHttpService()

	res, err := httpService.Post(http_protocol.HTTPRequest{
		URL:  "https://69fc-186-193-220-142.ngrok.io", // TODO: We must get the url from an env variable in config file
		Body: body,
	})
	fmt.Println(res) //TODO: Removes after
	if err != nil {
		return 0, err
	}

	return 1, nil //TODO: We must return idCustomer from response body
}
