package service

import (
	"errors"
	"testing"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*

 */

func TestLogin(t *testing.T) {
	repo := new(mocks.AuthRepository)
	t.Run("Success Login", func(t *testing.T) {
		inputData := auth.Core{Email: "sheena@duck.com", Password: "loonatheworld"}
		returnData := auth.Core{ID: 1, Name: "Sheena", Email: "sheena@duck.com", Password: "$2a$10$pWC8x2deWxzX6TjovlONTuLwK.S0p1vyf006nfjO98ZpeZ9qdeisK", Role: "User"}
		repo.On("FindUser", mock.Anything).Return(returnData, nil).Once()
		srv := New(repo)
		response, _, err := srv.Login(inputData)
		assert.Nil(t, err)
		assert.Equal(t, returnData, response)
		repo.AssertExpectations(t)

	})

	t.Run("Failed Login, Wrong Password", func(t *testing.T) {
		repo.On("FindUser", mock.Anything).Return(auth.Core{}, nil, errors.New("failed to login, password didn't match, please check password again")).Once()
		srv := New(repo)
		response, _, err := srv.Login(auth.Core{Email: "sheena@duck.com", Password: "", Role: "User"})
		assert.NotNil(t, err)
		assert.Equal(t, response, auth.Core{})
		repo.AssertExpectations(t)
	})

	t.Run("Failed Login, There's empty input field", func(t *testing.T) {
		inputKosong := auth.Core{Email: "sheena@duck.com", Password: "loonatheworld"}
		repo.On("FindUser", mock.Anything).Return(auth.Core{}, nil, errors.New("failed to login, password didn't match, please check password again")).Once()
		srv := New(repo)
		_, _, err := srv.Login(inputKosong)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "failed to login, password didn't match, please check password again")
		repo.AssertExpectations(t)
	})

}
