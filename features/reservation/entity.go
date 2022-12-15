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
	Address       string
	PricePerNight int
	BookedStart   time.Time
	BookedEnd     time.Time
	Reservation   []ReservationCore
}

type Payment struct {
	CreditCard    string
	Name          string
	CardNumber    string
	Cvv           string
	Month         string
	Year          string
	ReservationID uint
}

type History struct {
	ID          uint
	Reservation ReservationData
	Homestay    HomestayData
}

type ReservationData struct {
	ID         uint
	HomestayID uint
	StartDate  time.Time
	EndDate    time.Time
	TotalPrice int
}

type HomestayData struct {
	ID      uint
	Name    string
	Address string
}

type ServiceInterface interface {
	CheckAvailability(input ReservationCore) (data Homestay, err error)
	CreatePayment(input ReservationCore) (err error)
	GetHistory(id uint) (data []History, err error)
}
type RepositoryInterface interface {
	CheckAvailability(input ReservationCore) (data Homestay, err error)
	CreatePayment(input ReservationCore) (err error)
	GetHistory(id uint) (data []History, err error)
}
