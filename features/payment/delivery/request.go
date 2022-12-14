package delivery

import (
	payment "github.com/GP-3-Kelompok-2/airbnb-app-project/features/payment"
)

type PaymentRequest struct {
	HomestayID int
	StartDate  string
	EndDate    string
	Duration   int
	TotalPrice int
	CreditCard string `json:"credit_card" form:"credit_card"`
	Name       string `json:"name" form:"name"`
	CardNumber string `json:"card_number" form:"card_number"`
	Cvv        string `json:"cvv" form:"cvv"`
	Month      string `json:"month" form:"month"`
	Year       string `json:"year" form:"year"`
}

func requestToCore(PaymentInput PaymentRequest) payment.FirstCore {

	return payment.FirstCore{
		ReservationRequest: payment.ReservationRequest{
			HomestayID: PaymentInput.HomestayID,
			StartDate:  PaymentInput.StartDate,
			EndDate:    PaymentInput.EndDate,
		},
		HomestayResponse: payment.HomestayResponse{
			Duration:   PaymentInput.Duration,
			TotalPrice: PaymentInput.TotalPrice,
		},
		PaymentCore: payment.PaymentCore{
			CreditCard: PaymentInput.CreditCard,
			Name:       PaymentInput.Name,
			CardNumber: PaymentInput.CardNumber,
			Cvv:        PaymentInput.Cvv,
			Month:      PaymentInput.Month,
			Year:       PaymentInput.Year,
		},
	}
}
