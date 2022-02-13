package service

import (
	"fmt"
	"net/http"

	"github.com/yantoledo/input-service/entity/customer"
	"github.com/yantoledo/input-service/infra/http_protocol"
)

type CustomerService struct {
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (c *CustomerService) Insert(customer *customer.Customer) (int, error) {
	url := "https://7a3f-186-193-220-142.ngrok.io" // TODO: We must get the url from a env variable in config file

	httpService := http_protocol.NewHttpService()

	header := c.GetHeader()
	client := &http.Client{}
	response, err := httpService.Post(url, customer, header, client)
	fmt.Println(response) //TODO: Remove
	if err != nil {
		return 0, err
	}

	return 1, nil //TODO: We must return idCustomer from response body
}

func (c *CustomerService) GetHeader() http.Header {
	key := "Content_type"
	value := []string{"application/json"}
	header := http.Header(map[string][]string{key: value})

	return header
}
