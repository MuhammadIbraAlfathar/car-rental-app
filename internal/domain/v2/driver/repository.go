package driverV2

import (
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	FindById(driverId int) (*v2.Driver, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindById(driverId int) (*v2.Driver, error) {
	var driver *v2.Driver
	if err := r.db.First(&driver, driverId).Error; err != nil {
		return nil, err
	}

	return driver, nil
}
