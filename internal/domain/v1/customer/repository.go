package customerV1

import (
	"errors"
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	"gorm.io/gorm"
)

type Repository interface {
	Create(customer *v1Schema.Customer) (*v1Schema.Customer, error)
	GetAll() ([]*v1Schema.Customer, error)
	FindById(customerId int) (*v1Schema.Customer, error)
	Update(customer *v1Schema.Customer) (*v1Schema.Customer, error)
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

func (r *repository) GetAll() ([]*v1Schema.Customer, error) {
	var customer []*v1Schema.Customer
	if err := r.db.Find(&customer).Error; err != nil {
		return nil, errors.New("no data in record")
	}

	return customer, nil
}

func (r *repository) FindById(customerId int) (*v1Schema.Customer, error) {
	var customer *v1Schema.Customer
	if err := r.db.First(&customer, customerId).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *repository) Update(customer *v1Schema.Customer) (*v1Schema.Customer, error) {
	if err := r.db.Save(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}
