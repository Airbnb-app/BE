package delivery

import (
	// "errors"
	// "log"
	"net/http"

	// "strconv"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"

	"github.com/labstack/echo/v4"
)

type ReservationDelivery struct {
	reservationService reservation.ServiceInterface
}

func New(service reservation.ServiceInterface, e *echo.Echo) {
	handler := &ReservationDelivery{
		reservationService: service,
	}

	e.POST("/reservations/check", handler.CheckAvailability, middlewares.JWTMiddleware())
}

func (d *ReservationDelivery) CheckAvailability(c echo.Context) error {
	input := ReservationRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}
	dataInput := ToCore(input)
	res, err := d.reservationService.CheckAvailability(dataInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCore(res)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("available reservation", dataResponse))
}
