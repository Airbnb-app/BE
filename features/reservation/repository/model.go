package repository

import (
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	StartDate  time.Time
	EndDate    time.Time
	Duration   int
	TotalPrice int
	Homestay   Homestay
	UserID     uint
	HomestayID uint
	Payment    Payment
}

type User struct {
	gorm.Model
	Name        string
	Reservation []Reservation
}

type Homestay struct {
	gorm.Model
	Name          string
	PricePerNight int
	BookedStart   time.Time
	BookedEnd     time.Time
	Reservation   []Reservation
}

type Payment struct {
	gorm.Model
	CreditCard    string
	Name          string
	CardNumber    string
	Cvv           string
	Month         string
	Year          string
	ReservationID uint
}

func fromCore(dataCore reservation.ReservationCore) Reservation {
	reservationGorm := Reservation{
		StartDate:  dataCore.StartDate,
		EndDate:    dataCore.EndDate,
		Duration:   dataCore.Duration,
		TotalPrice: dataCore.TotalPrice,
		UserID:     dataCore.UserID,
		HomestayID: dataCore.HomestayID,
	}
	return reservationGorm
}

func (dataModel *Reservation) toCore() reservation.ReservationCore {
	return reservation.ReservationCore{
		ID:         dataModel.ID,
		StartDate:  dataModel.StartDate,
		EndDate:    dataModel.EndDate,
		Duration:   dataModel.Duration,
		TotalPrice: dataModel.Duration * dataModel.Homestay.PricePerNight,
		UserID:     dataModel.UserID,
		HomestayID: dataModel.HomestayID,
		Homestay: reservation.Homestay{
			ID:            dataModel.Homestay.ID,
			Name:          dataModel.Homestay.Name,
			PricePerNight: dataModel.Homestay.PricePerNight,
			BookedStart:   dataModel.Homestay.BookedStart,
			BookedEnd:     dataModel.Homestay.BookedEnd,
		},
		Payment: reservation.Payment{
			ID:            dataModel.Payment.ID,
			CreditCard:    dataModel.Payment.CreditCard,
			Name:          dataModel.Payment.Name,
			CardNumber:    dataModel.Payment.CardNumber,
			Cvv:           dataModel.Payment.Cvv,
			Month:         dataModel.Payment.Month,
			Year:          dataModel.Payment.Year,
			ReservationID: dataModel.Payment.ReservationID,
		},
	}
}

func (dataModel *Homestay) toCore() reservation.Homestay {
	return reservation.Homestay{
		ID:            dataModel.ID,
		PricePerNight: dataModel.PricePerNight,
		BookedStart:   dataModel.BookedStart,
		BookedEnd:     dataModel.BookedEnd,
	}
}

func toCoreList(dataModel []Reservation) []reservation.ReservationCore {
	var dataCore []reservation.ReservationCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
