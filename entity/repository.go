package entity

type CustomerRepository interface {
	Insert(Name string, UniqueID, UniqueClientID, Source int) error
}
