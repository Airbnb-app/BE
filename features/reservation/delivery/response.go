package delivery

import (
	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
)

type ReservationResponse struct {
	ID         uint
	Duration   int
	TotalPrice int
}

func fromCore(dataCore reservation.ReservationCore) ReservationResponse {
	return ReservationResponse{
		ID:         dataCore.ID,
		Duration:   dataCore.Duration,
		TotalPrice: dataCore.TotalPrice,
	}
}
