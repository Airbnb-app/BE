package mysql

import (
	"fmt"
	"log"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/config"
	// homestayRepo "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	// migrateDB(db)

	return db
}

func MigrateDB(db *gorm.DB) {
	// db.AutoMigrate(&userRepo.User{})
	// db.AutoMigrate(&homestayRepo.Homestay{})
	// db.AutoMigrate(&menteeRepo.Mentee{})
	// db.AutoMigrate(&logRepo.Log{})
}
