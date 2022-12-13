package delivery

import (
	"errors"
	"log"
	"net/http"

	"strconv"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
	"github.com/labstack/echo/v4"
)

type HomestayDelivery struct {
	homestayService homestay.ServiceInterface
}

func New(service homestay.ServiceInterface, e *echo.Echo) {
	handler := &HomestayDelivery{
		homestayService: service,
	}

	e.POST("/homestays", handler.CreateHomestay, middlewares.JWTMiddleware())
	e.GET("homestays/:id", handler.GetHomestayById, middlewares.JWTMiddleware())
	e.GET("/homestays", handler.GetAllHomestays, middlewares.JWTMiddleware())
	e.PUT("/homestays/:id", handler.UpdateHomestay, middlewares.JWTMiddleware())
	e.DELETE("/homestays/:id", handler.DeleteHomestay, middlewares.JWTMiddleware())

}

func (d *HomestayDelivery) CreateHomestay(c echo.Context) error {
	roleToken := middlewares.ExtractTokenUserRole(c)
	if roleToken != "Hoster" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Data only can be added by Hoster"))
	}
	dataInput := HomestayRequest{}
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	userId := middlewares.ExtractTokenUserId(c)
	dataInput.UserID = uint(userId)

	image1, _ := c.FormFile("image1")
	if image1 != nil {
		urlImage1, err := helper.UploadImage(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage1)
		dataInput.Image1 = urlImage1
	} else {
		dataInput.Image1 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	image2, _ := c.FormFile("image2")
	if image2 != nil {
		urlImage2, err := helper.UploadImage(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage2)
		dataInput.Image2 = urlImage2
	} else {
		dataInput.Image2 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	image3, _ := c.FormFile("image3")
	if image3 != nil {
		urlImage3, err := helper.UploadImage(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage3)
		dataInput.Image3 = urlImage3
	} else {
		dataInput.Image3 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	dataCore := requestToCore(dataInput)
	err := d.homestayService.CreateHomestay(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success Create New Homestay"))
}

func (d *HomestayDelivery) GetHomestayById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	result, err := d.homestayService.GetHomestayById(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))

	}
	dataResponse := fromCoreDetail(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get Homestay by Id", dataResponse))
}

func (d *HomestayDelivery) GetAllHomestays(c echo.Context) error {
	keyword := c.QueryParam("name")
	helper.LogDebug("\n search keyword: ", keyword)

	res, err := d.homestayService.GetAllHomestays(keyword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success get all homestays", dataResponse))
}

func (d *HomestayDelivery) UpdateHomestay(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	dataInput := HomestayRequest{}
	errBind := c.Bind(&dataInput)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}

	image1, _ := c.FormFile("image1")
	if image1 != nil {
		urlImage1, err := helper.UploadImage(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage1)
		dataInput.Image1 = urlImage1
	} else {
		dataInput.Image1 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	image2, _ := c.FormFile("image2")
	if image2 != nil {
		urlImage2, err := helper.UploadImage(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage2)
		dataInput.Image2 = urlImage2
	} else {
		dataInput.Image2 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	image3, _ := c.FormFile("image3")
	if image3 != nil {
		urlImage3, err := helper.UploadImage(c)
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage3)
		dataInput.Image2 = urlImage3
	} else {
		dataInput.Image2 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	dataUpdateCore := requestToCore(dataInput)
	_, err := d.homestayService.UpdateHomestay(dataUpdateCore, uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success update homestays", dataUpdateCore))

}

func (d *HomestayDelivery) DeleteHomestay(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	err := d.homestayService.DeleteHomestay(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Success Delete Homestay",
	})
}
