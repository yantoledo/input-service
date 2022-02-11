package process_customer

type CustomerDtoInput struct {
	Name           string
	UniqueID       int
	UniqueClientID int
	Source         int
}

type CustomerDtoOutput struct {
	UniqueID       int
	UniqueClientID int
}
