package service

type CustomerServiceOutput struct {
	IdCustomer     int    `json:"unique_id"`
	Name           string `json:"name"`
	UniqueClientID int    `json:"unique_client_id"`
}
