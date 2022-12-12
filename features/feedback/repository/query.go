package repository

import (
	"errors"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback"
	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.RepositoryInterface {
	return &feedbackRepository{
		db: db,
	}
}

// CreateFeedback implements feedback.RepositoryInterface
func (r *feedbackRepository) CreateFeedback(input feedback.FeedbackCore) (err error) {
	feedbackGorm := fromCore(input)
	tx := r.db.Create(&feedbackGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}
