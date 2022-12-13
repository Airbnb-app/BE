package repository

import (
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	Name          string
	Address       string
	Image1        string
	Image2        string
	Image3        string
	Description   string
	PricePerNight int
	BookedStart   time.Time
	BookedEnd     time.Time
	UserId        uint
	User          User
}

type HomestayForUser struct {
	ID            uint
	Name          string
	Address       string
	Image1        string
	Description   string
	PricePerNight int
}

type User struct {
	gorm.Model
	Name     string
	Homestay []Homestay
}

func fromCore(dataCore homestay.HomestayCore) Homestay {
	userGorm := Homestay{
		Name:          dataCore.Name,
		Address:       dataCore.Address,
		Image1:        dataCore.Image1,
		Image2:        dataCore.Image2,
		Image3:        dataCore.Image3,
		Description:   dataCore.Description,
		PricePerNight: dataCore.PricePerNight,
		UserId:        dataCore.UserId,
	}
	return userGorm
}

func (dataModel *Homestay) toCore() homestay.HomestayCore {
	return homestay.HomestayCore{
		ID:            dataModel.ID,
		Name:          dataModel.Name,
		Address:       dataModel.Address,
		Image1:        dataModel.Image1,
		Image2:        dataModel.Image2,
		Image3:        dataModel.Image3,
		Description:   dataModel.Description,
		PricePerNight: dataModel.PricePerNight,
		User: homestay.User{
			ID:   dataModel.User.ID,
			Name: dataModel.User.Name,
		},
	}
}

func toCoreList(dataModel []Homestay) []homestay.HomestayCore {
	var dataCore []homestay.HomestayCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
