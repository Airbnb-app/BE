package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"

type InsertRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Image1   string `json:"image1" form:"image1"`
}

type UpdateRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func toCore(i interface{}) user.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return user.Core{
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
			Role:     cnv.Role,
			Image1:   cnv.Image1,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return user.Core{
			Name:     cnv.Name,
			Email:    cnv.Email,
			Password: cnv.Password,
			Role:     cnv.Role,
		}
	}

	return user.Core{}
}
