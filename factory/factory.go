package factory

import (
	userDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/delivery"
	userRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/repository"
	userService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/service"

	authDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/delivery"
	authRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/repository"
	authService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)
}
