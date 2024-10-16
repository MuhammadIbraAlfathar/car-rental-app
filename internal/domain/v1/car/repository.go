package carV1

import (
	"errors"
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	"gorm.io/gorm"
)

type Repository interface {
	Create(car *v1Schema.Car) (*v1Schema.Car, error)
	GetAll() ([]*v1Schema.Car, error)
	FindById(carId int) (*v1Schema.Car, error)
	Update(car *v1Schema.Car) (*v1Schema.Car, error)
	Delete(car *v1Schema.Car) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(car *v1Schema.Car) (*v1Schema.Car, error) {
	err := r.db.Create(&car).Error
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repository) GetAll() ([]*v1Schema.Car, error) {
	var car []*v1Schema.Car
	if err := r.db.Find(&car).Error; err != nil {
		return nil, errors.New("no data in record")
	}

	return car, nil
}

func (r *repository) FindById(carId int) (*v1Schema.Car, error) {
	var car *v1Schema.Car
	if err := r.db.First(&car, carId).Error; err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repository) Update(car *v1Schema.Car) (*v1Schema.Car, error) {
	if err := r.db.Save(car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

func (r *repository) Delete(car *v1Schema.Car) error {
	return r.db.Delete(car).Error
}
