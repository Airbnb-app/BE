package repository

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/payment"

type Payment struct {
	CreditCard string
	Name       string
	CardNumber string
	Cvv        string
	Month      string
	Year       string
}

func fromCore(dataCore payment.FirstCore) Payment {
	paymentGorm := Payment{
		CreditCard: dataCore.PaymentCore.CreditCard,
		Name:       dataCore.PaymentCore.Name,
		CardNumber: dataCore.PaymentCore.CardNumber,
		Cvv:        dataCore.PaymentCore.Cvv,
		Month:      dataCore.PaymentCore.Month,
		Year:       dataCore.PaymentCore.Year,
	}
	return paymentGorm
}
