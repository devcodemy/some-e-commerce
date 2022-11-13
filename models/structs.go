package models

type AbstractUser struct {
	EmailAddress string
	PhoneNumber  string
	PasswordHash string
}

type Currency struct {
	Name         string
	Code         string
	ExchangeRate float32 // exhange rate to local currency
}

type ProductType struct {
	Name        string
	Description string
	Type        interface{}
}

func (c Currency) Exchange(amount float32) float32 {
	return amount * c.ExchangeRate
}
