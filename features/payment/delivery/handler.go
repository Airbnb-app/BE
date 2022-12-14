package delivery

import (
	"net/http"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/payment"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
	"github.com/labstack/echo/v4"
)

type PaymentDelivery struct {
	paymentService payment.ServiceInterface
}

func New(service payment.ServiceInterface, e *echo.Echo) {
	handler := &PaymentDelivery{
		paymentService: service,
	}

	e.POST("/reservations", handler.CreatePayment, middlewares.JWTMiddleware())

}

func (delivery *PaymentDelivery) CreatePayment(c echo.Context) error {

	paymentInput := PaymentRequest{}

	errBind := c.Bind(&paymentInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Process Failed. Please check your input "+errBind.Error()))
	}

	paymentCore := requestToCore(paymentInput)

	err := delivery.paymentService.CreatePayment(paymentCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success reservation, see you later"))
}
