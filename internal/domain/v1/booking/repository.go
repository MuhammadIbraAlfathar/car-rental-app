package bookingV1

import (
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	"gorm.io/gorm"
)

type Repository interface {
	Create(booking *v1Schema.Booking) (*v1Schema.Booking, error)
	GetAll() ([]*v1Schema.Booking, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(booking *v1Schema.Booking) (*v1Schema.Booking, error) {
	err := r.db.Create(&booking).Error
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repository) GetAll() ([]*v1Schema.Booking, error) {
	var booking []*v1Schema.Booking
	if err := r.db.Preload("Customer").Preload("Car").Find(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}
