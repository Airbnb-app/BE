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
func (r *reservationRepository) CheckAvailability(input reservation.ReservationCore) (data reservation.Homestay, err error) {
	var homestay Homestay
	tx := r.db.Not("booked_start BETWEEN ? AND ? AND booked_end BETWEEN ? AND ?", input.StartDate, input.EndDate, input.StartDate, input.EndDate).First(&homestay, input.HomestayID)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = homestay.toCore()
	return data, nil
}
