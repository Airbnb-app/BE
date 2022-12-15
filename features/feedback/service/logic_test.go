package service

import (
	"errors"
	"testing"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateFeedback(t *testing.T) {
	repo := new(mocks.FeedbackRepository)
	t.Run("Success Create feedback", func(t *testing.T) {
		inputData := feedback.FeedbackCore{Rating: "*****", Feedback: "lorem ipsum blablabla", UserId: 1, UserName: "alta", HomestayID: 1}
		repo.On("CreateFeedback", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.CreateFeedback(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create feedback", func(t *testing.T) {
		inputData := feedback.FeedbackCore{Rating: "*****", Feedback: "lorem ipsum blablabla", UserId: 1, UserName: "alta", HomestayID: 1}
		repo.On("CreateFeedback", inputData).Return(errors.New("failed create log, error query")).Once()
		srv := New(repo)
		err := srv.CreateFeedback(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed create log, error query", err.Error())
		repo.AssertExpectations(t)
	})
}
