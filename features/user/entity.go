package user

import (
	"time"

	homestay "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
)

type Core struct {
	ID       uint
	Name     string `valiidate:"required"`
	Email    string `valiidate:"required,email,unique"`
	Password string `valiidate:"required"`
	Role     string `validiate:"required"`
	Image    struct {
		Image1 string
		Image2 string
		Image3 string
	}
	CreatedAt time.Time
	UpdatedAt time.Time
	Homestay  []homestay.HomestayCore
}

type ServiceInterface interface {
	Create(input Core) error
	Update(input Core, id uint) error
	Get(id uint) (data Core, err error)
	Upgrade(input Core, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	Create(input Core) error
	Update(input Core, id uint) error
	Get(id uint) (data Core, err error)
	Upgrade(input Core, id uint) error
	Delete(id uint) error
	FindUser(email string) (data Core, err error)
}
