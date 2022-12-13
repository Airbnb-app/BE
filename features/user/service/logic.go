package service

import (
	"errors"
	"strings"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.Core) (err error) {
	input.Role = "User"
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik
	_, errFindEmail := service.userRepository.FindUser(input.Email)

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return errFindEmail
	}

	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return err
	}

	input.Password = string(bytePass)

	errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return err
	}

	return nil
}

func (service *userService) Get(id uint) (data user.Core, err error) {
	data, err = service.userRepository.Get(id)
	if err != nil {
		log.Error(err.Error())
		return user.Core{}, err
	}
	return data, err

}

func (service *userService) Update(input user.Core, id uint) error {
	if input.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		input.Password = string(generate)
	}
	err := service.userRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (service *userService) Delete(id uint) error {
	err := service.userRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (s *userService) Upgrade(data user.Core, id uint) error {
	errUpgrade := s.userRepository.Upgrade(data, id)
	if errUpgrade != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}
