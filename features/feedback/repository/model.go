package repository

import (
	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Rating     string
	Feedback   string
	UserId     uint
	UserName   string
	HomestayID uint
	// User       User
}

type User struct {
	gorm.Model
	Feedback []Feedback `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Homestay struct {
	gorm.Model
	Name     string
	Feedback []Feedback `gorm:"foreignKey:HomestayID"`
}

func fromCore(dataCore feedback.FeedbackCore) Feedback {
	feedbackGorm := Feedback{
		Rating:     dataCore.Rating,
		Feedback:   dataCore.Feedback,
		UserId:     dataCore.UserId,
		UserName:   dataCore.UserName,
		HomestayID: dataCore.HomestayID,
	}
	return feedbackGorm
}
