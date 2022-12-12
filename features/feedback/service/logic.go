package service

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback"
)

type feedbackService struct {
	feedbackRepository feedback.RepositoryInterface
}

func New(repo feedback.RepositoryInterface) feedback.ServiceInterface {
	return &feedbackService{
		feedbackRepository: repo,
	}
}

// CreateFeedback implements feedback.ServiceInterface
func (s *feedbackService) CreateFeedback(input feedback.FeedbackCore) (err error) {
	errCreate := s.feedbackRepository.CreateFeedback(input)
	if errCreate != nil {
		return errors.New("failed create log, error query")
	}

	return nil
}
