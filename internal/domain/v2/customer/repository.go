package customerV2

import (
	"errors"
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/gorm"
)

type Repository interface {
	Create(customer *v2.CustomerNew) (*v2.CustomerNew, error)
	GetAll() ([]*v2.CustomerNew, error)
	FindById(customerId int) (*v2.CustomerNew, error)
	Update(customer *v2.CustomerNew) (*v2.CustomerNew, error)
	Delete(customerId int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(customer *v2.CustomerNew) (*v2.CustomerNew, error) {
	err := r.db.Create(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *repository) GetAll() ([]*v2.CustomerNew, error) {
	var customer []*v2.CustomerNew
	if err := r.db.Find(&customer).Error; err != nil {
		return nil, errors.New("no data in record")
	}

	return customer, nil
}

func (r *repository) FindById(customerId int) (*v2.CustomerNew, error) {
	var customer *v2.CustomerNew
	if err := r.db.First(&customer, customerId).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *repository) Update(customer *v2.CustomerNew) (*v2.CustomerNew, error) {
	if err := r.db.Save(customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func (r *repository) Delete(customerId int) error {
	return r.db.Where("id = ?", customerId).Delete(&v2.CustomerNew{}).Error
}
