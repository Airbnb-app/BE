package repository

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/payment"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) payment.RepositoryInterface {
	return &paymentRepository{
		db: db,
	}
}

// CreateFeedback implements feedback.RepositoryInterface
func (r *paymentRepository) CreatePayment(input payment.FirstCore) error {
	paymentGorm := fromCore(input)
	tx := r.db.Create(&paymentGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}
