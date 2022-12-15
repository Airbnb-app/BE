package service

import (
	"errors"
	"testing"
	"time"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCheckAvailability(t *testing.T) {
	repo := new(mocks.ReservationRepository)
	t.Run("Success CheckAvailability", func(t *testing.T) {
		inputData := reservation.ReservationCore{StartDate: time.Now(), EndDate: time.Now(), HomestayID: 1}
		repo.On("CheckAvailability", inputData).Return(reservation.Homestay{}, nil).Once()
		srv := New(repo)
		res, err := srv.CheckAvailability(inputData)
		assert.NoError(t, err)
		assert.Equal(t, res, reservation.Homestay{})
		repo.AssertExpectations(t)
	})

	t.Run("Failed CheckAvailability", func(t *testing.T) {
		inputData := reservation.ReservationCore{StartDate: time.Now(), EndDate: time.Now(), HomestayID: 1}
		repo.On("CheckAvailability", inputData).Return(reservation.Homestay{}, errors.New("failed get data, error query")).Once()
		srv := New(repo)
		_, err := srv.CheckAvailability(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed get data, error query", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestCreatePayment(t *testing.T) {
	repo := new(mocks.ReservationRepository)
	inputData := reservation.ReservationCore{StartDate: time.Now(), EndDate: time.Now(), HomestayID: 1, Duration: 1, TotalPrice: 1, UserID: 1}
	t.Run("success CreatePayment", func(t *testing.T) {
		repo.On("CreatePayment", inputData).Return(nil).Once()
		srv := New(repo)
		err := srv.CreatePayment(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed CreatePayment", func(t *testing.T) {
		repo.On("CreatePayment", inputData).Return(errors.New("failed create payment, error query")).Once()
		srv := New(repo)
		err := srv.CreatePayment(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed create payment, error query", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetHistory(t *testing.T) {
	repo := new(mocks.ReservationRepository)
	var UserId uint
	returnData := []reservation.ReservationCore{{ID: 1, HomestayID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalPrice: 1, Duration: 1, UserID: 1}}
	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetHistory", UserId).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetHistory(UserId)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].HomestayID, response[0].HomestayID)
		repo.AssertExpectations(t)

	})

	t.Run("failed Get Data", func(t *testing.T) {
		repo.On("GetHistory", UserId).Return(nil, errors.New("failed get data, error query")).Once()
		srv := New(repo)
		response, err := srv.GetHistory(UserId)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "failed get data, error query", err.Error())
		repo.AssertExpectations(t)

	})
}
