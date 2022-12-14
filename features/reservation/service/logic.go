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
