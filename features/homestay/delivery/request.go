package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"

type HomestayRequest struct {
	Name          string `json:"name" form:"name"`
	Address       string `json:"address" form:"address"`
	Image1        string `json:"image1" form:"image1"`
	Image2        string `json:"image2" form:"image2"`
	Image3        string `json:"image3" form:"image3"`
	Description   string `json:"description" form:"description"`
	PricePerNight int    `json:"price_per_night" form:"price_per_night"`
	UserID        uint   `json:"user_id" form:"user_id"`
}

func requestToCore(homestayInput HomestayRequest) homestay.HomestayCore {
	homestayCoreData := homestay.HomestayCore{
		Name:          homestayInput.Name,
		Address:       homestayInput.Address,
		Image1:        homestayInput.Image1,
		Image2:        homestayInput.Image2,
		Image3:        homestayInput.Image3,
		Description:   homestayInput.Description,
		PricePerNight: homestayInput.PricePerNight,
		UserID:        homestayInput.UserID,
	}
	return homestayCoreData
}
