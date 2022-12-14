package delivery

import (
	"math"
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
)

// ID         uint
// 	StartDate  time.Time
// 	EndDate    time.Time
// 	Duration   int
// 	TotalPrice int
// 	Homestay   Homestay
// 	UserID     uint
// 	HomestayID uint

type ReservationRequest struct {
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	HomestayID int    `json:"homestay_id" form:"homestay_id"`
}

var dateLayout = "2006-01-02"

func ToCore(reservationInput ReservationRequest) reservation.ReservationCore {
	start, _ := time.Parse(dateLayout, reservationInput.StartDate)
	end, _ := time.Parse(dateLayout, reservationInput.EndDate)
	period := int(math.Ceil(end.Sub(start).Hours() / 24))
	return reservation.ReservationCore{
		StartDate: start,
		EndDate:   end,
		Duration:  period,
	}
}
