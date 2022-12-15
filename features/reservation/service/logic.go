package service

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
	// "github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	// "github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
)

type reservationService struct {
	reservationRepo reservation.RepositoryInterface
}

func New(data reservation.RepositoryInterface) reservation.ServiceInterface {
	return &reservationService{
		reservationRepo: data,
	}
}

// CheckAvailability implements reservation.ServiceInterface
func (s *reservationService) CheckAvailability(input reservation.ReservationCore) (data reservation.Homestay, err error) {
	data, err = s.reservationRepo.CheckAvailability(input)
	if err != nil {
		return data, errors.New("failed get data, error query")

	}
	return data, nil
}

// CreatePayment implements reservation.ServiceInterface
func (s *reservationService) CreatePayment(input reservation.ReservationCore) (err error) {
	err = s.reservationRepo.CreatePayment(input)
	if err != nil {
		return errors.New("failed create payment, error query")
	}

	return nil
}

// GetHistory implements reservation.ServiceInterface
func (s *reservationService) GetHistory(id uint) (data []reservation.History, err error) {
	data, err = s.reservationRepo.GetHistory(id)
	if err != nil {
		return data, errors.New("failed get data, error query")
	}
	return data, nil
}
