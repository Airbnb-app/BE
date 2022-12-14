package reservation

import "time"

type ReservationCore struct {
	ID         uint
	StartDate  time.Time
	EndDate    time.Time
	Duration   int
	TotalPrice int
	Homestay   Homestay
	UserID     uint
	HomestayID uint
	Payment    Payment
}

type User struct {
	ID          uint
	Name        string
	Reservation []ReservationCore
}

type Homestay struct {
	ID            uint
	Name          string
	PricePerNight int
	BookedStart   time.Time
	BookedEnd     time.Time
	Reservation   []ReservationCore
}

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

type ServiceInterface interface {
	CheckAvailability(input ReservationCore) (data Homestay, err error)
	CreatePayment(input ReservationCore) (err error)
}
type RepositoryInterface interface {
	CheckAvailability(input ReservationCore) (data Homestay, err error)
	CreatePayment(input ReservationCore) (err error)
}
