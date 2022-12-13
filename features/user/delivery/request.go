package delivery

import "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"

type InsertRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Image    struct {
		Image1 string `json:"image1" form:"image1"`
		Image2 string `json:"image2" form:"image2"`
		Image3 string `json:"image3" form:"image3"`
	}
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
			Image: struct {
				Image1 string
				Image2 string
				Image3 string
			}(cnv.Image),
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
