package bookingV2

import (
	"errors"
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	Create(booking *v2.BookingNew) (*v2.BookingNew, error)
	GetAll() ([]*v2.BookingNew, error)
	FindByCustomerId(customerId int) ([]*v2.BookingNew, error)
	FindById(bookId int) (*v2.BookingNew, error)
	Update(booking *v2.BookingNew) (*v2.BookingNew, error)
	Delete(booking *v2.BookingNew) error
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

func (r *repository) GetAll() ([]*v2.BookingNew, error) {
	var booking []*v2.BookingNew
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}
	if err := r.db.Preload("Customer").Preload("Car").Find(&booking).Error; err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *repository) FindByCustomerId(customerId int) ([]*v2.BookingNew, error) {
	var booking []*v2.BookingNew
	if err := r.db.Preload("Customer").Preload("Car").Where("customer_id = ?", customerId).Find(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repository) FindById(bookId int) (*v2.BookingNew, error) {
	var booking *v2.BookingNew
	if err := r.db.Preload("Customer").Preload("Car").First(&booking, bookId).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repository) Update(booking *v2.BookingNew) (*v2.BookingNew, error) {
	if err := r.db.Save(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func (r *repository) Delete(booking *v2.BookingNew) error {
	return r.db.Delete(booking).Error
}
