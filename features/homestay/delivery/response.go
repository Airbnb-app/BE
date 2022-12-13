package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"

type HomestayResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Image1        string `json:"image1"`
	Image2        string `json:"image2"`
	Image3        string `json:"image3"`
	Description   string `json:"description"`
	PricePerNight int    `json:"price_per_night"`
	UserID        uint   `json:"user_id"`
	Owner         string `json:"owner"`
}

type HomestayDetailResponse struct {
	ID            uint               `json:"id"`
	Name          string             `json:"name"`
	Address       string             `json:"address"`
	Image1        string             `json:"image1"`
	Image2        string             `json:"image2"`
	Image3        string             `json:"image3"`
	Description   string             `json:"description"`
	PricePerNight int                `json:"price_per_night"`
	UserID        uint               `json:"user_id"`
	Owner         string             `json:"owner"`
	Feedback      []FeedbackResponse `json:"feedback"`
}

type FeedbackResponse struct {
	ID         uint   `json:"id"`
	Rating     string `json:"rating"`
	Feedback   string `json:"feedback"`
	Poster     string `json:"poster"`
	HomestayID uint   `json:"homestay_id"`
}

func fromCore(dataCore homestay.HomestayCore) HomestayResponse {
	return HomestayResponse{
		ID:            dataCore.ID,
		Name:          dataCore.Name,
		Address:       dataCore.Address,
		Image1:        dataCore.Image1,
		Image2:        dataCore.Image2,
		Image3:        dataCore.Image3,
		Description:   dataCore.Description,
		PricePerNight: dataCore.PricePerNight,
		UserID:        dataCore.UserID,
		Owner:         dataCore.User.Name,
	}
}

func fromCoreDetail(dataCore homestay.HomestayCore) HomestayDetailResponse {
	var arrFeedbacks []FeedbackResponse

	for _, v := range dataCore.Feedback {
		arrFeedbacks = append(arrFeedbacks, FeedbackResponse{
			ID:         v.ID,
			Rating:     v.Rating,
			Feedback:   v.Feedback,
			Poster:     v.UserName,
			HomestayID: v.HomestayID,
		})
	}
	return HomestayDetailResponse{
		ID:            dataCore.ID,
		Name:          dataCore.Name,
		Address:       dataCore.Address,
		Image1:        dataCore.Image1,
		Image2:        dataCore.Image2,
		Image3:        dataCore.Image3,
		Description:   dataCore.Description,
		PricePerNight: dataCore.PricePerNight,
		UserID:        dataCore.UserID,
		Owner:         dataCore.User.Name,
		Feedback:      arrFeedbacks,
	}
}

func fromCoreList(dataCore []homestay.HomestayCore) []HomestayResponse {
	var dataResponse []HomestayResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
