package delivery

import (
	"time"

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

type HistoryResponse struct {
	HomestayID uint      `json:"homestay_id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalPrice int       `json:"total_price"`
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

func fromCoreHistory(dataCore reservation.History) HistoryResponse {
	return HistoryResponse{
		HomestayID: dataCore.Reservation.HomestayID,
		Name:       dataCore.Homestay.Name,
		Address:    dataCore.Homestay.Address,
		StartDate:  dataCore.Reservation.StartDate,
		EndDate:    dataCore.Reservation.EndDate,
		TotalPrice: dataCore.Reservation.TotalPrice,
	}
}

func fromCoreList(dataCore []reservation.History) []HistoryResponse {
	var dataResponse []HistoryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreHistory(v))
	}
	return dataResponse
}
