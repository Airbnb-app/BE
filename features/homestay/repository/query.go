package repository

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	"gorm.io/gorm"
)

type homestayRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.RepositoryInterface {
	return &homestayRepository{
		db: db,
	}
}

// InsertHomestay implements homestay.RepositoryInterface
func (r *homestayRepository) InsertHomestay(data homestay.HomestayCore) (row int, err error) {
	homestayGorm := fromCore(data)
	tx := r.db.Create(&homestayGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetHomestayById implements homestay.RepositoryInterface
func (r *homestayRepository) GetHomestayById(id uint) (data homestay.HomestayCore, err error) {
	var homestay Homestay

	tx := r.db.First(&homestay, id)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = homestay.toCore()
	return data, nil
}

// GetAllHomestays implements homestay.RepositoryInterface
func (r *homestayRepository) GetAllHomestays(keyword string) (data []homestay.HomestayCore, err error) {
	var homestays []Homestay
	if keyword == "" {
		tx := r.db.Find(&homestays)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := r.db.Where("name LIKE ?", "%"+keyword+"%").Find(&homestays)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	dataCore := toCoreList(homestays)
	return dataCore, nil
}

// UpdateHomestay implements homestay.RepositoryInterface
func (r *homestayRepository) UpdateHomestay(input homestay.HomestayCore, id uint) (data homestay.HomestayCore, err error) {
	var homestay Homestay

	inputData := fromCore(input)
	tx := r.db.Model(&homestay).Where("id = ?", id).Updates(inputData)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = homestay.toCore()
	return data, nil
}

// DeleteHomestay implements homestay.RepositoryInterface
func (r *homestayRepository) DeleteHomestay(id uint) (err error) {
	var homestay Homestay

	tx := r.db.Delete(&homestay, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}
