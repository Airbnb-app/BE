package service

import (
	"errors"
	"testing"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	repo := new(mocks.UserRepository)
	returnData := user.Core{ID: 1, Name: "Myoboku", Email: "alta", Role: "User", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg"}
	t.Run("Success Get Data", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		response, err := srv.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		repo.AssertExpectations(t)

	})

	t.Run("Failed Get data", func(t *testing.T) {
		repo.On("Get", mock.Anything).Return(user.Core{}, errors.New("failed to get data")).Once()
		srv := New(repo)
		response, err := srv.Get(2)
		assert.NotNil(t, err)
		assert.Equal(t, response, user.Core{})
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepository)
	id := uint(1)
	t.Run("Success Delete Data", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(id)
		assert.Nil(t, err)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete Data", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("")).Once()
		srv := New(repo)
		err := srv.Delete(uint(0))
		assert.NotNil(t, err)
		assert.Equal(t, "", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success Update Data", func(t *testing.T) {
		var idParam uint = 1
		inputData := user.Core{Name: "Myoboku", Email: "alta", Role: "User", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg"}
		repo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(inputData, idParam)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update Data", func(t *testing.T) {
		var idParam uint = 1
		inputData := user.Core{Name: "Myoboku", Email: "alta", Role: "User", Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg"}
		repo.On("Update", mock.Anything, mock.Anything).Return(errors.New("")).Once()
		srv := New(repo)
		err := srv.Update(inputData, idParam)
		assert.NotNil(t, err)
		assert.Equal(t, "", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestUpgrade(t *testing.T) {
	repo := new(mocks.UserRepository)
	t.Run("Success Upgrade Data", func(t *testing.T) {
		var idParam uint = 1
		inputData := user.Core{Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg"}
		repo.On("Upgrade", mock.Anything, mock.Anything).Return(nil).Once()
		srv := New(repo)
		err := srv.Upgrade(inputData, idParam)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed Upgrade Data", func(t *testing.T) {
		var idParam uint = 1
		inputData := user.Core{Image1: "https://s3address/imgegi/aogalkw-image1.jpg", Image2: "https://s3address/imgegi/aogalkw-image2.jpg", Image3: "https://s3address/imgegi/aogalkw-image3.jpg"}
		repo.On("Upgrade", mock.Anything, mock.Anything).Return(errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.Upgrade(inputData, idParam)
		assert.NotNil(t, err)
		assert.Equal(t, "failed to insert data, error query", err.Error())
		repo.AssertExpectations(t)
	})
}
