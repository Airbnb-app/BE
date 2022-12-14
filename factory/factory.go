package factory

import (
	userDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/delivery"
	userRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/repository"
	userService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/service"

	authDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/delivery"
	authRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/repository"
	authService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/service"

	homestayDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/delivery"
	homestayRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/repository"
	homestayService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/service"

	feedbackDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/delivery"
	feedbackRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/repository"
	feedbackService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/service"

	reservationDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation/delivery"
	reservationRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation/repository"
	reservationService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation/service"

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

	homestayRepoFactory := homestayRepo.New(db)
	homestayServiceFactory := homestayService.New(homestayRepoFactory)
	homestayDelivery.New(homestayServiceFactory, e)

	feedbackRepoFactory := feedbackRepo.New(db)
	feedbackServiceFactory := feedbackService.New(feedbackRepoFactory)
	feedbackDelivery.New(feedbackServiceFactory, e)

	reservationRepoFactory := reservationRepo.New(db)
	reservationServiceFactory := reservationService.New(reservationRepoFactory)
	reservationDelivery.New(reservationServiceFactory, e)
}
