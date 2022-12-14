package payment

import "time"

type Payment struct {
	ID            uint
	CreditCard    string
	Name          string
	CardNumber    string
	Cvv           string
	Month         string
	Year          string
	ReservationID uint
}

type Reservation struct {
	ID        uint
	StartDate time.Time
	EndDate   time.Time
	Payment   Payment
}
