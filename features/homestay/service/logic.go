package service

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	// "github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	// "github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
)

type homestayService struct {
	homestayRepo homestay.RepositoryInterface
}

func New(data homestay.RepositoryInterface) homestay.ServiceInterface {
	return &homestayService{
		homestayRepo: data,
	}
}

// CreateHomestay implements homestay.ServiceInterface
func (s *homestayService) CreateHomestay(data homestay.HomestayCore) (err error) {
	_, errCreate := s.homestayRepo.InsertHomestay(data)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetHomestayById implements homestay.ServiceInterface
func (s *homestayService) GetHomestayById(id uint) (data homestay.HomestayCore, err error) {
	data, err = s.homestayRepo.GetHomestayById(id)
	if err != nil {
		return data, errors.New("failed get data, error query")

	}
	return
}

// GetAllHomestays implements homestay.ServiceInterface
func (s *homestayService) GetAllHomestays(keyword string) (data []homestay.HomestayCore, err error) {
	data, err = s.homestayRepo.GetAllHomestays(keyword)
	if err != nil {
		return nil, err
	}
	return
}

// UpdateHomestay implements homestay.ServiceInterface
func (s *homestayService) UpdateHomestay(input homestay.HomestayCore, id uint) (data homestay.HomestayCore, err error) {
	data, err = s.homestayRepo.UpdateHomestay(input, id)
	if err != nil {
		return homestay.HomestayCore{}, errors.New("failed update data, error query")
	}
	return data, nil
}

// DeleteHomestay implements homestay.ServiceInterface
func (s *homestayService) DeleteHomestay(id uint) (err error) {
	err = s.homestayRepo.DeleteHomestay(id)
	if err != nil {
		return errors.New("failed to delete data, error query")
	}
	return nil
}
