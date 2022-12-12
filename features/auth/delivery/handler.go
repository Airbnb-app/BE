package delivery

import (
	"net/http"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/utils/helper"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService auth.ServiceInterface
}

func New(service auth.ServiceInterface, e *echo.Echo) {
	handler := &AuthHandler{
		authService: service,
	}
	e.POST("/login", handler.Login)
}

func (handler *AuthHandler) Login(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 	"message": "failed to get bind data",
		// })
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("failed to bind data"))
	}

	dataCore := ToCore(userInput)
	result, token, err := handler.authService.Login(dataCore)

	if err != nil {
		// return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		// 	"message": "failed to get token data" + err.Error(),
		// })
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed to Login. "+err.Error()))
	}
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success",
	// 	"name":    name,
	// 	"token":   result,
	// })
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Login Success.", FromCore(result, token)))
}
