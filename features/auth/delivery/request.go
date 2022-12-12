package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth"

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq UserRequest) auth.Core {
	userCore := auth.Core{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
