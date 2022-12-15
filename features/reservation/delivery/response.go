package delivery

import (
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
)

type ReservationResponse struct {
	ID         uint `json:"id"`
	Duration   int  `json:"duration"`
	TotalPrice int  `json:"total_price"`
}

type HomestayResponse struct {
	HomestayID    uint `json:"homestay_id"`
	PricePerNight int  `json:"price_per_night"`
	Duration      int  `json:"duration"`
	TotalPrice    int  `json:"total_price"`
}

type HomestayTrip struct {
	Name    string `json:"homestay_name"`
	Address string `json:"address"`
}
type HistoryResponse struct {
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Duration   int       `json:"duration"`
	TotalPrice int       `json:"total_price"`
	// Homestay   HomestayTrip
	HomestayName    string `json:"homestay_name"`
	HomestayAddress string `json:"homestay_address"`
	HomestayID      uint   `json:"homestay_id"`
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

func fromCoreTrip(dataCore reservation.ReservationCore) HistoryResponse {
	return HistoryResponse{
		StartDate:  dataCore.StartDate,
		EndDate:    dataCore.EndDate,
		Duration:   dataCore.Duration,
		TotalPrice: dataCore.TotalPrice,
		HomestayID: dataCore.HomestayID,
		// Homestay: HomestayTrip{
		// 	Name:    dataCore.Homestay.Name,
		// 	Address: dataCore.Homestay.Address,
		// },
		HomestayName: dataCore.Homestay.Name,
	}
}

func TripArr(dataCore []reservation.ReservationCore) []HistoryResponse {
	var dataResponse []HistoryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreTrip(v))
	}
	return dataResponse
}
