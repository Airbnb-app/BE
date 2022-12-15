package payment

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	CreditCard    string
	Name          string
	CardNumber    string
	Cvv           string
	Month         string
	Year          string
	ReservationID uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Reservation struct {
	gorm.Model
	StartDate time.Time
	EndDate   time.Time
	Payment   Payment `gorm:"foreignKey:ReservationID"`
}
