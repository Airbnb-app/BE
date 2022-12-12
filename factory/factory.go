package factory

import (
<<<<<<< HEAD
	userDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/delivery"
	userRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/repository"
	userService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/user/service"

	authDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/delivery"
	authRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/repository"
	authService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/auth/service"
=======
	homestayDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/delivery"
	homestayRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/repository"
	homestayService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/service"

	feedbackDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/delivery"
	feedbackRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/repository"
	feedbackService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/service"
>>>>>>> 4baaf9103d16b9cb49854ca54ddb5d9a64d83119

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
<<<<<<< HEAD
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)
=======
	homestayRepoFactory := homestayRepo.New(db)
	homestayServiceFactory := homestayService.New(homestayRepoFactory)
	homestayDelivery.New(homestayServiceFactory, e)

	feedbackRepoFactory := feedbackRepo.New(db)
	feedbackServiceFactory := feedbackService.New(feedbackRepoFactory)
	feedbackDelivery.New(feedbackServiceFactory, e)
>>>>>>> 4baaf9103d16b9cb49854ca54ddb5d9a64d83119
}
