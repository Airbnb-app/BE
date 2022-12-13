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

// doesn't need
// type Homestay struct {
// 	ID       uint
// 	Name     string
// 	Feedback []Feedback
// }

// type User struct {
// 	ID       uint
// 	Name     string
// 	Feedback []Feedback
// }

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
