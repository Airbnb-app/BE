package payment

type FirstCore struct {
	ReservationRequest ReservationRequest
	HomestayResponse   HomestayResponse
	PaymentCore        PaymentCore
}

type PaymentCore struct {
	CreditCard string
	Name       string
	CardNumber string
	Cvv        string
	Month      string
	Year       string
}

type ReservationRequest struct {
	HomestayID int
	StartDate  string
	EndDate    string
}

type HomestayResponse struct {
	Duration   int
	TotalPrice int
}

type ServiceInterface interface {
	CreatePayment(data FirstCore) error
}

type RepositoryInterface interface {
	CreatePayment(data FirstCore) error
}
