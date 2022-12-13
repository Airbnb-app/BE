package delivery

import (
	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback"
)

type FeedbackRequest struct {
	UserID     uint   `json:"user_id" form:"user_id"`
	UserName   string `json:"user_name" form:"user_name"`
	HomestayID uint   `json:"homestay_id" form:"homestay_id"`
	Rating     string `json:"rating" form:"rating"`
	Feedback   string `json:"feedback" form:"feedback"`
}

func requestToCore(FeedbackInput FeedbackRequest) feedback.FeedbackCore {
	return feedback.FeedbackCore{
		UserId:     FeedbackInput.UserID,
		UserName:   FeedbackInput.UserName,
		HomestayID: FeedbackInput.HomestayID,
		Rating:     FeedbackInput.Rating,
		Feedback:   FeedbackInput.Feedback,
	}
}
