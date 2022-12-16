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
	Image1   string
	Image2   string
	Image3   string
	Homestay []Homestay `gorm:"constraint:OnDelete:CASCADE;"`
	Feedback []Feedback `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Homestay struct {
	gorm.Model
	Name          string
	Address       string
	Image1        string
	Description   string
	PricePerNight int
	UserID        uint
}

type Feedback struct {
	gorm.Model
	Rating     string
	Feedback   string
	UserId     uint
	UserName   string
	HomestayID uint
	// User       User
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
		Image1:   dataCore.Image1,
		Image2:   dataCore.Image2,
		Image3:   dataCore.Image3,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() user.Core {
	var arrHomestay []user.Homestay
	for _, val := range dataModel.Homestay {
		arrHomestay = append(arrHomestay, user.Homestay{
			ID:            val.ID,
			Name:          val.Name,
			Address:       val.Address,
			Image1:        val.Image1,
			Description:   val.Description,
			PricePerNight: val.PricePerNight,
		})
	}
	return user.Core{
		ID:       dataModel.ID,
		Name:     dataModel.Name,
		Email:    dataModel.Email,
		Password: dataModel.Password,
		Role:     dataModel.Role,
		Image1:   dataModel.Image1,
		Image2:   dataModel.Image2,
		Image3:   dataModel.Image3,
		Homestay: arrHomestay,
	}
}

// mengubah slice struct model gorm ke slice struct core
// func toCoreList(dataModel []User) []_user.Core {
// 	var dataCore []_user.Core
// 	for _, v := range dataModel {
// 		dataCore = append(dataCore, v.toCore())
// 	}
// 	return dataCore
// }
