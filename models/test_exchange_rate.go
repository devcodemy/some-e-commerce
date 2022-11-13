package models

import "log"

func TestExchange() {
	uah := Currency{
		Name:         "Gryvna",
		Code:         "UAH",
		ExchangeRate: 1.0,
	}

	usd := Currency{
		Name:         "Dollar",
		Code:         "USD",
		ExchangeRate: 41.0,
	}

	eur := Currency{
		Name:         "Euro",
		Code:         "EUR",
		ExchangeRate: 40.15,
	}

	var am float32 = 99.0
	log.Printf("Exchange operation: %v %v Equal to %v %v", am, uah.Code, uah.Exchange(am), uah.Code)
	log.Printf("Exchange operation: %v %v Equal to %v %v", am, usd.Code, usd.Exchange(am), uah.Code)
	log.Printf("Exchange operation: %v %v Equal to %v %v", am, eur.Code, eur.Exchange(am), uah.Code)

	log.Println("---")

	var i float32 = 1.0
	var s float32 = 0.11
	for ; i < 10.0; i += s {
		log.Printf("Exchange operation: %v EUR Equal to %v UAH", i, eur.Exchange(i))
	}
} // float32 is anougth for Currency operations!
