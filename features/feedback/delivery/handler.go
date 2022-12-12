package delivery

import (
	"net/http"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
	"github.com/labstack/echo/v4"
)

type FeedbackDelivery struct {
	feedbackService feedback.ServiceInterface
}

func New(service feedback.ServiceInterface, e *echo.Echo) {
	handler := &FeedbackDelivery{
		feedbackService: service,
	}
	e.POST("/feedbacks", handler.CreateFeedback, middlewares.JWTMiddleware())
}

func (delivery *FeedbackDelivery) CreateFeedback(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	name := middlewares.ExtractTokenUserName(c)
	helper.LogDebug("\n extracttokenname= ", name)

	if userId == 0 {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Process Failed. Please check your input"))
	}

	feedbackInput := FeedbackRequest{}

	errBind := c.Bind(&feedbackInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Process Failed. Please check your input "+errBind.Error()))
	}

	feedbackInput.UserName = name
	feedbackInput.UserID = uint(userId)

	feedbackCore := requestToCore(feedbackInput)

	err := delivery.feedbackService.CreateFeedback(feedbackCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("Success add feedback"))
}
