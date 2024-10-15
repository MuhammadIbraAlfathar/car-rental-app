package driverIncentiveV2

import (
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	Create(driverIncentive *v2.DriverIncentive) (*v2.DriverIncentive, error)
	FindByBookingId(bookingId int) (*v2.DriverIncentive, error)
	Update(driverIncentive *v2.DriverIncentive) (*v2.DriverIncentive, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(driverIncentive *v2.DriverIncentive) (*v2.DriverIncentive, error) {
	err := r.db.Create(&driverIncentive).Error
	if err != nil {
		return nil, err
	}

	return driverIncentive, nil
}

func (r *repository) FindByBookingId(bookingId int) (*v2.DriverIncentive, error) {
	var driverIncentive *v2.DriverIncentive
	if err := r.db.Where("booking_id = ?", bookingId).Find(&driverIncentive).Error; err != nil {
		return nil, err
	}

	return driverIncentive, nil
}

func (r *repository) Update(driverIncentive *v2.DriverIncentive) (*v2.DriverIncentive, error) {
	if err := r.db.Save(&driverIncentive).Error; err != nil {
		return nil, err
	}

	return driverIncentive, nil
}
