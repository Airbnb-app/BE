package repository

import (
	_homestay "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	homestay "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/repository"
	_user "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"
	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	Name     string
	Email    string `validate:"required,email"`
	Password string `valudate:"required"`
	Role     string `valudate:"required"`
	Image    struct {
		Image1 string
		Image2 string
		Image3 string
	}
	Homestay []homestay.Homestay
	Feedback []Feedback
}

type Feedback struct {
	gorm.Model
	Feedback string
	UserID   uint
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _user.Core) User {
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Role:     dataCore.Role,
	}
	return userGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *User) toCore() _user.Core {
	var arrHomestay []_homestay.HomestayCore
	for _, val := range dataModel.Homestay {
		arrHomestay = append(arrHomestay, _homestay.HomestayCore{
			ID:            val.ID,
			Name:          val.Name,
			Address:       val.Address,
			Image1:        val.Image1,
			Description:   val.Description,
			PricePerNight: val.PricePerNight,
		})
	}
	return _user.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		Homestay:  arrHomestay,
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
