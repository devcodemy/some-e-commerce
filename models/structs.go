package models

type AbstractUser struct {
	ID           string
	EmailAddress string
	PhoneNumber  string
	PasswordHash string
}

type Currency struct {
	ID     string
	Name   string
	Code   string
	Symbol string
}

type ExchangeRate struct {
	ID          string
	Name        string
	Description string
	Current     Currency
	Convert     Currency
	Rate        float32
}
type ProductType struct {
	ID          string
	Name        string
	Description string
	Type        interface{}
}

// ---
