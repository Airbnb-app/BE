package repository

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *userRepository) Create(input user.Core) error {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// GetById implements user.RepositoryInterface
func (repo *userRepository) Get(id uint) (data user.Core, err error) {
	var user User

	tx := repo.db.Preload("Homestay").First(&user, id)

	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = user.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *userRepository) Update(input user.Core, id uint) error {
	userGorm := fromCore(input)
	var user User
	tx := repo.db.Model(&user).Where("id = ?", id).Updates(&userGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed")
	}
	return nil
}

// Delete implements user.Repository
func (repo *userRepository) Delete(id uint) error {
	var user User
	tx := repo.db.Unscoped().Delete(&user, id) // proses delete
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (repo *userRepository) FindUser(email string) (result user.Core, err error) {
	var userData User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	result = userData.toCore()

	return result, nil
}

func (repo *userRepository) Upgrade(input user.Core, id uint) error {
	userGorm := fromCore(input)
	tx := repo.db.Where("id = ?", id).Updates(&userGorm) // proses update
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert image1 failed")
	}

	return nil
}
