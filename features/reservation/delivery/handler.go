package delivery

import (
	// "errors"
	// "log"
	"math"
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
	e.POST("/reservations", handler.Payment, middlewares.JWTMiddleware())
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

	start := dataInput.StartDate
	end := dataInput.EndDate
	period := int(math.Ceil(end.Sub(start).Hours() / 24))

	dataResponse := fromCoreAvail(res)
	dataResponse.Duration = period
	dataResponse.TotalPrice = dataInput.Duration * dataResponse.PricePerNight

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("available reservation", dataResponse))
}

func (d *ReservationDelivery) Payment(c echo.Context) error {
	input := PaymentRequest{}
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}
	dataInput := ToCorePayment(input)
	err := d.reservationService.CreatePayment(dataInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Success reservation, see you later"))
}
