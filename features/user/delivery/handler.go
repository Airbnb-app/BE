package delivery

import (
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
	e.GET("/users", handler.Get, middlewares.JWTMiddleware(), middlewares.UserOnlySameId)
	// e.POST("/users/upgrade", handler.Upgrade, middlewares.JWTMiddleware())
	e.PUT("/users", handler.Update, middlewares.JWTMiddleware(), middlewares.UserOnlySameId)
	e.DELETE("/users", handler.Delete, middlewares.JWTMiddleware(), middlewares.UserOnlySameId)

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

	results, err := delivery.userService.Get()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	dataResponse := fromCore(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success read user.", dataResponse))
}

func (delivery *UserDelivery) Update(c echo.Context) error {
	userInput := UpdateRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data. "+errBind.Error()))
	}

	dataCore := toCore(userInput)
	err := delivery.userService.Update(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed update data. "+err.Error()))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success update data."))
}

func (delivery *UserDelivery) Delete(c echo.Context) error {
	err := delivery.userService.Delete()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Success delete data."))
}
