package customer_service

type CustomerProcessed struct {
	IdCustomer     int    `json:"unique_id"`
	Name           string `json:"name"`
	UniqueClientID int    `json:"unique_client_id"`
	Source         int    `json:"source"`
}
