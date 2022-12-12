package factory

import (
	homestayDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/delivery"
	homestayRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/repository"
	homestayService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/service"

	feedbackDelivery "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/delivery"
	feedbackRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/repository"
	feedbackService "github.com/GP-3-Kelompok-2/airbnb-app-project/features/feedback/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	homestayRepoFactory := homestayRepo.New(db)
	homestayServiceFactory := homestayService.New(homestayRepoFactory)
	homestayDelivery.New(homestayServiceFactory, e)

	feedbackRepoFactory := feedbackRepo.New(db)
	feedbackServiceFactory := feedbackService.New(feedbackRepoFactory)
	feedbackDelivery.New(feedbackServiceFactory, e)
}
