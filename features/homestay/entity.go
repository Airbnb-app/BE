package homestay

import "time"

type HomestayCore struct {
	ID            uint
	Name          string
	Address       string
	Image1        string
	Image2        string
	Image3        string
	Description   string
	PricePerNight int
	BookedStart   time.Time
	BookedEnd     time.Time
	UserID        uint
	User          User
	Feedback      []Feedback
}

type User struct {
	ID   uint
	Name string
}

type Feedback struct {
	ID         uint
	Rating     string
	Feedback   string
	UserName   string
	HomestayID uint
}

type ServiceInterface interface {
	CreateHomestay(data HomestayCore) (err error)
	GetAllHomestays(keyword string) (data []HomestayCore, err error)
	GetHomestayById(id uint) (data HomestayCore, err error)
	UpdateHomestay(input HomestayCore, id uint) (data HomestayCore, err error)
	DeleteHomestay(id uint) (err error)
}

type RepositoryInterface interface {
	InsertHomestay(data HomestayCore) (row int, err error)
	GetAllHomestays(keyword string) (data []HomestayCore, err error)
	GetHomestayById(id uint) (data HomestayCore, err error)
	UpdateHomestay(input HomestayCore, id uint) (data HomestayCore, err error)
	DeleteHomestay(id uint) (err error)
}
