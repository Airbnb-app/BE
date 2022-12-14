package delivery

import (
	"math"
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
)

type ReservationRequest struct {
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	HomestayID int    `json:"homestay_id" form:"homestay_id"`
	Homestay   Homestay
}

type PaymentRequest struct {
	StartDate  string `json:"start_date" form:"start_date"`
	EndDate    string `json:"end_date" form:"end_date"`
	HomestayID int    `json:"homestay_id" form:"homestay_id"`
	Duration   int    `json:"duration" form:"duration"`
	UserID     uint   `json:"user_id" form:"user_id"`
	Payment    Payment
}

type Homestay struct {
	ID          uint
	BookedStart time.Time
	BookedEnd   time.Time
}

type Payment struct {
	CreditCard string `json:"credit_card" form:"credit_card"`
	Name       string `json:"name" form:"name"`
	CardNumber string `json:"credit_number" form:"credit_number"`
	Cvv        string `json:"cvv" form:"cvv"`
	Month      string `json:"month" form:"month"`
	Year       string `json:"year" form:"year"`
}

var dateLayout = "2006-01-02"

func ToCore(reservationInput ReservationRequest) reservation.ReservationCore {
	start, _ := time.Parse(dateLayout, reservationInput.StartDate)
	end, _ := time.Parse(dateLayout, reservationInput.EndDate)
	period := int(math.Ceil(end.Sub(start).Hours() / 24))
	return reservation.ReservationCore{
		StartDate:  start,
		EndDate:    end,
		Duration:   period,
		HomestayID: uint(reservationInput.HomestayID),
		Homestay: reservation.Homestay{
			ID:          reservationInput.Homestay.ID,
			BookedStart: reservationInput.Homestay.BookedStart,
			BookedEnd:   reservationInput.Homestay.BookedEnd,
		},
	}
}

func ToCorePayment(reservationInput PaymentRequest) reservation.ReservationCore {
	start, _ := time.Parse(dateLayout, reservationInput.StartDate)
	end, _ := time.Parse(dateLayout, reservationInput.EndDate)
	period := int(math.Ceil(end.Sub(start).Hours() / 24))
	return reservation.ReservationCore{
		StartDate:  start,
		EndDate:    end,
		Duration:   period,
		UserID:     reservationInput.UserID,
		HomestayID: uint(reservationInput.HomestayID),
		Payment: reservation.Payment{
			CreditCard: reservationInput.Payment.CreditCard,
			Name:       reservationInput.Payment.Name,
			CardNumber: reservationInput.Payment.CardNumber,
			Cvv:        reservationInput.Payment.Cvv,
			Month:      reservationInput.Payment.Month,
			Year:       reservationInput.Payment.Year,
		},
	}
}
