package usecase

type CustomerDtoInput struct {
	UniqueID       int `json:"unique_id"`
	UniqueClientID int `json:"unique_client_id"`
	Source         int `json:"source"`
}

type CustomerDtoOutput struct {
	IdCustomer     int    `json:"unique_id"`
	Name           string `json:"name"`
	UniqueClientID int    `json:"unique_client_id"`
}
