package payment

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CreditCard    string
	Name          string
	CardNumber    string
	Cvv           string
	Month         string
	Year          string
	ReservationID uint
}

type Reservation struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
	Payment   Payment
}
