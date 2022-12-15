package service

import (
	"errors"
	"testing"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateHomestay(t *testing.T) {
	repo := new(mocks.HomestayRepository)
	t.Run("Success Create User", func(t *testing.T) {
		inputData := homestay.HomestayCore{Name: "Myoboku", Address: "Tokyo", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}
		repo.On("InsertHomestay", inputData).Return(1, nil).Once()
		srv := New(repo)
		err := srv.CreateHomestay(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create Homestay, duplicate entry", func(t *testing.T) {
		inputData := homestay.HomestayCore{Name: "Myoboku", Address: "Tokyo", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}
		repo.On("InsertHomestay", inputData).Return(0, errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.CreateHomestay(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to insert data, error query", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetAllHomestays(t *testing.T) {
	repo := new(mocks.HomestayRepository)
	keyword := "Myobo"
	returnData := []homestay.HomestayCore{{ID: 1, Name: "Myoboku", Address: "Tokyo", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}}
	t.Run("Success Get All Homestay", func(t *testing.T) {
		repo.On("GetAllHomestays", keyword).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetAllHomestays(keyword)
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All Homestay", func(t *testing.T) {
		repo.On("GetAllHomestays", keyword).Return(nil, errors.New("failed get data, error query")).Once()
		srv := New(repo)
		response, err := srv.GetAllHomestays(keyword)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetHomestayById(t *testing.T) {
	repo := new(mocks.HomestayRepository)
	returnData := homestay.HomestayCore{ID: 1, Name: "Myoboku", Address: "Tokyo", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}
	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("GetHomestayById", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.GetHomestayById(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		repo.AssertExpectations(t)

	})

	t.Run("Failed Get data", func(t *testing.T) {
		repo.On("GetHomestayById", mock.Anything).Return(homestay.HomestayCore{}, errors.New("failed to get data")).Once()
		srv := New(repo)
		response, err := srv.GetHomestayById(2)
		assert.NotNil(t, err)
		assert.Equal(t, response, homestay.HomestayCore{})
		repo.AssertExpectations(t)
	})
}

func TestDeleteHomestay(t *testing.T) {
	repo := new(mocks.HomestayRepository)
	id := uint(1)
	t.Run("Success Delete Data", func(t *testing.T) {
		repo.On("DeleteHomestay", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.DeleteHomestay(id)
		assert.Nil(t, err)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete Data", func(t *testing.T) {
		repo.On("DeleteHomestay", mock.Anything).Return(errors.New("failed to delete data, error query")).Once()
		srv := New(repo)
		err := srv.DeleteHomestay(uint(0))
		assert.NotNil(t, err)
		assert.Equal(t, "failed to delete data, error query", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdateHomestay(t *testing.T) {
	repo := new(mocks.HomestayRepository)
	t.Run("Success Update Data", func(t *testing.T) {
		var idParam uint = 1
		returnData := homestay.HomestayCore{ID: 1, Name: "Myoboku", Address: "Tokyo-to", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}
		inputData := homestay.HomestayCore{Name: "Myoboku", Address: "Tokyo-to", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}
		repo.On("UpdateHomestay", mock.Anything, mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.UpdateHomestay(inputData, idParam)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update Data", func(t *testing.T) {
		var idParam uint = 1
		inputData := homestay.HomestayCore{Name: "Myoboku", Address: "Tokyo-to", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg", Description: "High establishment you never see", PricePerNight: 2000, UserID: 1}
		repo.On("UpdateHomestay", mock.Anything, mock.Anything).Return(homestay.HomestayCore{}, errors.New("failed update data, error query")).Once()
		srv := New(repo)
		_, err := srv.UpdateHomestay(inputData, idParam)
		assert.NotNil(t, err)
		assert.Equal(t, "failed update data, error query", err.Error())
		repo.AssertExpectations(t)
	})
}
