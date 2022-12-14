package repository

import (
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
func (r *reservationRepository) CheckAvailability(input reservation.ReservationCore) (data reservation.ReservationCore, err error) {
	var reservation Reservation

	tx := r.db.Preload("Homestay").Not("homestays.booked_start BETWEEN ? AND ? AND homestays.booked_end BETWEEN ? AND ?", input.StartDate, input.EndDate, input.StartDate, input.EndDate).First(&reservation)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = reservation.toCore()
	return data, nil
}
