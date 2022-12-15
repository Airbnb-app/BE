package repository

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.RepositoryInterface {
	return &reservationRepository{
		db: db,
	}
}

// CheckAvailability implements reservation.RepositoryInterface
func (r *reservationRepository) CheckAvailability(input reservation.ReservationCore) (data reservation.Homestay, err error) {
	var homestay Homestay
	tx := r.db.Not("booked_start BETWEEN ? AND ? AND booked_end BETWEEN ? AND ?", input.StartDate, input.EndDate, input.StartDate, input.EndDate).First(&homestay, input.HomestayID)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = homestay.toCore()
	return data, nil
}

// CreatePayment implements reservation.RepositoryInterface
func (r *reservationRepository) CreatePayment(input reservation.ReservationCore) (err error) {
	var homestay Homestay
	inputGorm := fromCore(input)
	tx := r.db.Create(&Reservation{StartDate: inputGorm.StartDate,
		EndDate:    inputGorm.EndDate,
		Duration:   inputGorm.Duration,
		UserID:     inputGorm.UserID,
		HomestayID: inputGorm.HomestayID,
		TotalPrice: inputGorm.TotalPrice,
		Payment: Payment{CreditCard: inputGorm.Payment.CreditCard,
			Name:       inputGorm.Payment.Name,
			CardNumber: inputGorm.Payment.CardNumber,
			Cvv:        inputGorm.Payment.Cvv,
			Month:      inputGorm.Payment.Month,
			Year:       inputGorm.Payment.Year}})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	tx2 := r.db.Model(&homestay).Where("id = ?", input.HomestayID).Updates(&Homestay{BookedStart: input.StartDate, BookedEnd: input.EndDate})
	if tx2.Error != nil {
		return tx2.Error
	}
	return nil
}

// GetHistory implements reservation.RepositoryInterface
func (r *reservationRepository) GetHistory(UserId uint) (data []reservation.ReservationCore, err error) {
	var reservation []Reservation
	tx := r.db.Preload("Homestay").Where("user_id = ?", UserId).Find(&reservation)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := toCoreList(reservation)
	return dataCore, nil

}
