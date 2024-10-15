package driver_incentiveV2

import (
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	Create(driverIncentive *v2.DriverIncentive) (*v2.DriverIncentive, error)
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
