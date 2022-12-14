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

type ServiceInterface interface {
	CheckAvailability(input ReservationCore) (data ReservationCore, err error)
	// CreateHomestay(data HomestayCore) (err error)
	// GetAllHomestays(keyword string) (data []HomestayCore, err error)
	// GetHomestayById(id uint) (data HomestayCore, err error)
	// UpdateHomestay(input HomestayCore, id uint) (data HomestayCore, err error)
	// DeleteHomestay(id uint) (err error)
}
type RepositoryInterface interface {
	CheckAvailability(input ReservationCore) (data ReservationCore, err error)
	// InsertHomestay(data HomestayCore) (row int, err error)
	// GetAllHomestays(keyword string) (data []HomestayCore, err error)
	// GetHomestayById(id uint) (data HomestayCore, err error)
	// UpdateHomestay(input HomestayCore, id uint) (data HomestayCore, err error)
	// DeleteHomestay(id uint) (err error)
}
