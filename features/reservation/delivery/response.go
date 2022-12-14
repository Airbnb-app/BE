package delivery

import (
	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
)

type ReservationResponse struct {
	ID         uint
	Duration   int
	TotalPrice int
}

type HomestayResponse struct {
	HomestayID    uint `json:"homestay_id"`
	PricePerNight int  `json:"price_per_night"`
	Duration      int  `json:"duration"`
	TotalPrice    int  `json:"total_price"`
}

func fromCore(dataCore reservation.ReservationCore) ReservationResponse {
	return ReservationResponse{
		ID:         dataCore.ID,
		Duration:   dataCore.Duration,
		TotalPrice: dataCore.TotalPrice,
	}
}

func fromCoreAvail(dataCore reservation.Homestay) HomestayResponse {
	return HomestayResponse{
		HomestayID:    dataCore.ID,
		PricePerNight: dataCore.PricePerNight,
	}
}
