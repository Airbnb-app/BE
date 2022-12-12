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
}

type ServiceInterface interface {
	Create(input Core) error
	Update(input Core, id int) error
	GetById(id int) (data Core, err error)
	// Upgrade(input Core, id int) error
	Delete(id int) error
}

type RepositoryInterface interface {
	Create(input Core) error
	Update(input Core, id int) error
	GetById(id int) (data Core, err error)
	// Upgrade(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
}
