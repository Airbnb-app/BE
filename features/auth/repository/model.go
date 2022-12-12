package repository

import (
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

//DTO

func (data User) toCore() auth.Core {
	return auth.Core{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
