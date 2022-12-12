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
	UserId        uint   `json:"user_id"`
	Owner         string `json:"owner"`
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
		UserId:        dataCore.UserId,
		Owner:         dataCore.User.Name,
	}
}

func fromCoreList(dataCore []homestay.HomestayCore) []HomestayResponse {
	var dataResponse []HomestayResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
