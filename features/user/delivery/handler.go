package delivery

import (
	"errors"
	"log"
	"net/http"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/user"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}
	e.POST("/users", handler.Create)
	e.GET("/users", handler.Get, middlewares.JWTMiddleware())
	e.POST("/users/upgrade", handler.Upgrade, middlewares.JWTMiddleware())
	e.PUT("/users", handler.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", handler.Delete, middlewares.JWTMiddleware())

	//middlewares.UserOnlySameId = untuk membatasi akses user mengelola data diri sendiri saja
}

func (delivery *UserDelivery) Create(c echo.Context) error {
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.userService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success create data"))
}

func (delivery *UserDelivery) Get(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	x := uint(userId)
	results, err := delivery.userService.Get(x)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *UserDelivery) Update(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	x := uint(userId)
	userInput := UpdateRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.userService.Update(dataCore, x)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *UserDelivery) Delete(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	x := uint(userId)
	err := delivery.userService.Delete(x)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}

func (delivery *UserDelivery) Upgrade(c echo.Context) error {
	//upgrade cukup input request body image1
	userId := middlewares.ExtractTokenUserId(c)
	x := uint(userId)
	userInput := InsertRequest{}
	errBind := c.Bind(&userInput.Image) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	image1, _ := c.FormFile("image1")
	if image1 != nil {
		urlImage1, err := helper.UploadImage(c, "image1")
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage1)
		userInput.Image.Image1 = urlImage1
	} else {
		userInput.Image.Image1 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}
	image2, _ := c.FormFile("image2")
	if image2 != nil {
		urlImage2, err := helper.UploadImage(c, "image2")
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage2)
		userInput.Image.Image2 = urlImage2
	} else {
		userInput.Image.Image2 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	image3, _ := c.FormFile("image3")
	if image3 != nil {
		urlImage3, err := helper.UploadImage(c, "image3")
		if err != nil {
			return errors.New("registration failed. cannot upload data")
		}
		log.Print(urlImage3)
		userInput.Image.Image3 = urlImage3
	} else {
		userInput.Image.Image3 = "https://img1.wikia.nocookie.net/__cb20130610133347/onepiece/it/images/3/3d/Noland_bugiardo_2.png"
	}

	dataCore := toCore(userInput.Image)
	err := delivery.userService.Upgrade(dataCore, x)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data. "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success upgrade data"))
}
