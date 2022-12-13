package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Homestay []HomestayResponse
}

type HomestayResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Image1        string `json:"image1"`
	Description   string `json:"description"`
	PricePerNight int    `json:"price_per_night"`
}

func fromCore(dataCore user.Core) UserResponse {
	var arrHomestay []HomestayResponse
	for _, val := range dataCore.Homestay {
		arrHomestay = append(arrHomestay, HomestayResponse{
			ID:            val.ID,
			Name:          val.Name,
			Address:       val.Address,
			Image1:        val.Image1,
			Description:   val.Description,
			PricePerNight: val.PricePerNight,
		})
	}
	return UserResponse{
		ID:       dataCore.ID,
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Role:     dataCore.Role,
		Homestay: arrHomestay,
	}
}

// func fromCoreList(dataCore []user.Core) []UserResponse {
// 	var dataResponse []UserResponse
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCore(v))
// 	}
// 	return dataResponse
// }
