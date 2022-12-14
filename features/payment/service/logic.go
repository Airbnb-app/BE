package service

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/payment"
)

type paymentService struct {
	paymentRepository payment.RepositoryInterface
}

func New(repo payment.RepositoryInterface) payment.ServiceInterface {
	return &paymentService{
		paymentRepository: repo,
	}
}

// CreateFeedback implements feedback.ServiceInterface
func (s *paymentService) CreatePayment(input payment.FirstCore) error {
	errCreate := s.paymentRepository.CreatePayment(input)
	if errCreate != nil {
		return errors.New("failed add payment, error query")
	}

	return nil
}
