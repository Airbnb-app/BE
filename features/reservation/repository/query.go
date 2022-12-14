package repository

import (
	"github.com/GP-3-Kelompok-2/airbnb-app-project/features/reservation"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) reservation.RepositoryInterface {
	return &reservationRepository{
		db: db,
	}
}

// CheckAvailability implements reservation.RepositoryInterface
func (r *reservationRepository) CheckAvailability(input reservation.ReservationCore) (data reservation.ReservationCore, err error) {
	var reservation Reservation
	// idHome := input.HomestayID
	tx := r.db.Preload("Homestay").Not("Homestay.BookedStart BETWEEN ? AND ? AND Homestay.BookedEnd BETWEEN ? AND ?", input.StartDate, input.EndDate, input.StartDate, input.EndDate).First(&reservation)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = reservation.toCore()
	return data, nil
} /*
var properties []Property
	queryBuilder := fmt.Sprintf("SELECT * FROM bookings WHERE property_id = %d AND '%s' BETWEEN checkin_date AND checkout_date OR '%s' BETWEEN checkin_date AND checkout_date;", propertyId, checkinDate, checkoutData)
	// tx := repo.db.Raw(`
	// 	SELECT * FROM bookings WHERE property_id = @propertyID
	// 	AND @checkinDate BETWEEN checkin_date AND checkout_date
	// 	OR @checkoutDate BETWEEN checkin_date AND checkout_date;
	// `, sql.Named("propertyID", propertyId), sql.Named("checkinDate", checkinDate), sql.Named("checkoutDate", checkoutData)).Find(&properties)

	fmt.Println("\n\n query ", queryBuilder)

	tx := repo.db.Raw(queryBuilder).Find(&properties)
*/
