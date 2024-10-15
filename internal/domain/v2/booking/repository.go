package bookingV2

import (
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	Create(booking *v2.BookingNew) (*v2.BookingNew, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(booking *v2.BookingNew) (*v2.BookingNew, error) {
	err := r.db.Create(&booking).Error
	if err != nil {
		return nil, err
	}

	return booking, nil
}
