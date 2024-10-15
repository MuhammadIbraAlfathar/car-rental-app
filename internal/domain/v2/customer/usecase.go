package customerV2

import (
	"errors"
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
)

type UseCase struct {
	customerRepository Repository
}

func NewUseCase(customerRepository Repository) *UseCase {
	return &UseCase{
		customerRepository: customerRepository,
	}
}

func (uc *UseCase) CreateCustomer(req *CreateCustomerRequest) (*v2.CustomerNew, error) {
	customerEntity := v2.CustomerNew{
		Name:         req.Name,
		Nik:          req.Nik,
		PhoneNumber:  req.PhoneNumber,
		MembershipId: req.MembershipId,
	}

	newCustomer, err := uc.customerRepository.Create(&customerEntity)
	if err != nil {
		return nil, err
	}

	return newCustomer, nil
}

func (uc *UseCase) GetAllCustomer() ([]*v2.CustomerNew, error) {
	customer, err := uc.customerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (uc *UseCase) GetCustomerById(customerId int) (*v2.CustomerNew, error) {
	customer, err := uc.customerRepository.FindById(customerId)
	if err != nil {
		return nil, errors.New("customer not found")
	}

	return customer, nil

}
