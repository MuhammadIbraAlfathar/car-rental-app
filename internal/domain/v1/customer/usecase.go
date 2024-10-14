package customerV1

import (
	"errors"
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
)

type UseCase struct {
	customerRepository Repository
}

func NewUseCase(customerRepository Repository) *UseCase {
	return &UseCase{
		customerRepository: customerRepository,
	}
}

func (uc *UseCase) CreateCustomer(req *CreateCustomerRequest) (*v1Schema.Customer, error) {
	customerEntity := v1Schema.Customer{
		Name:        req.Name,
		Nik:         req.Nik,
		PhoneNumber: req.PhoneNumber,
	}

	newCustomer, err := uc.customerRepository.Create(&customerEntity)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (uc *UseCase) GetAllCustomer() ([]*v1Schema.Customer, error) {
	customer, err := uc.customerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (uc *UseCase) UpdateCustomer(customerId int, req *UpdatedCustomerRequest) (*v1Schema.Customer, error) {
	customer, err := uc.customerRepository.FindById(customerId)
	if err != nil {
		return nil, errors.New("customer not found")
	}
	customer.Name = req.Name
	customer.Nik = req.Nik
	customer.PhoneNumber = req.PhoneNumber

	updatedCustomer, err := uc.customerRepository.Update(customer)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (uc *UseCase) DeleteCustomer(customerId int) error {
	if err := uc.customerRepository.Delete(customerId); err != nil {
		return errors.New("user not found")
	}

	return nil
}
