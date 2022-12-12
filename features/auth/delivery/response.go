package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth"

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FromCore(data auth.Core, token string) UserResponse {
	return UserResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Role:  data.Role,
		Token: token,
	}
}
