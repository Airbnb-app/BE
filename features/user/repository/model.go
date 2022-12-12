package repository

import (
	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"
	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	Name     string
	Email    string `validate:"required,email"`
	Password string `valudate:"required"`
	Role     string `valudate:"required"`
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore user.Core) User {
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Role:     dataCore.Role,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() user.Core {
	return user.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []user.Core {
	var dataCore []user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
