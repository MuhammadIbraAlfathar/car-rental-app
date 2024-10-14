package customerV1

import (
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	"gorm.io/gorm"
)

type Repository interface {
	Create(customer *v1Schema.Customer) (*v1Schema.Customer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(customer *v1Schema.Customer) (*v1Schema.Customer, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}
