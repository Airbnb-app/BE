package user

import "time"

type Core struct {
	ID        uint
	Name      string `valiidate:"required"`
	Email     string `valiidate:"required,email,unique"`
	Password  string `valiidate:"required"`
	Role      string `validiate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Homestay  []Homestay
}

type Homestay struct {
	ID            uint
	Name          string
	Address       string
	Image1        string
	Description   string
	PricePerNight int
}

type ServiceInterface interface {
	Create(input Core) error
	Update(input Core) error
	Get() (data Core, err error)
	// Upgrade(input Core, id int) error
	Delete() error
}

type RepositoryInterface interface {
	Create(input Core) error
	Update(input Core) error
	Get() (data Core, err error)
	// Upgrade(input Core, id int) error
	Delete() error
	FindUser(email string) (data Core, err error)
}
