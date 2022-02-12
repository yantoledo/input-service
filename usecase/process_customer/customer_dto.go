package process_customer

type CustomerDtoInput struct {
	Name           string `json:"name"`
	UniqueID       int    `json:"unique_id"`
	UniqueClientID int    `json:"unique_client_id"`
	Source         int    `json:"source"`
}

type CustomerDtoOutput struct {
	idCustomer     int `json:"unique_id"`
	UniqueClientID int `json:"unique_client_id"`
}
