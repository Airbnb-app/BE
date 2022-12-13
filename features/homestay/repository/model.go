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
	UserID        uint
	User          User
	Feedback      []Feedback
}

type User struct {
	gorm.Model
	Name     string
	Homestay []Homestay
}

type Feedback struct {
	gorm.Model
	Rating     string
	Feedback   string
	UserName   string
	HomestayID uint
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
		UserID:        dataCore.UserID,
	}
	return userGorm
}

func (dataModel *Homestay) toCore() homestay.HomestayCore {
	var arrFeedbacks []homestay.Feedback
	for _, v := range dataModel.Feedback {
		arrFeedbacks = append(arrFeedbacks, homestay.Feedback{
			ID:         v.ID,
			Rating:     v.Rating,
			Feedback:   v.Feedback,
			UserName:   v.UserName,
			HomestayID: v.HomestayID,
		})
	}
	return homestay.HomestayCore{
		ID:            dataModel.ID,
		Name:          dataModel.Name,
		Address:       dataModel.Address,
		Image1:        dataModel.Image1,
		Image2:        dataModel.Image2,
		Image3:        dataModel.Image3,
		Description:   dataModel.Description,
		PricePerNight: dataModel.PricePerNight,
		UserID:        dataModel.UserID,
		User:          homestay.User{ID: dataModel.User.ID, Name: dataModel.User.Name},
		Feedback:      arrFeedbacks,
	}
}

func toCoreList(dataModel []Homestay) []homestay.HomestayCore {
	var dataCore []homestay.HomestayCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
