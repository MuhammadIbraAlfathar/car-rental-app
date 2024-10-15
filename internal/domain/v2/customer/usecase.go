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

func (uc *UseCase) UpdateCustomer(customerId int, req *UpdatedCustomerRequest) (*v2.CustomerNew, error) {
	customer, err := uc.customerRepository.FindById(customerId)
	if err != nil {
		return nil, errors.New("customer not found")
	}

	customerUpdate := &v2.CustomerNew{
		Id:           customer.Id,
		Name:         req.Name,
		Nik:          req.Nik,
		PhoneNumber:  req.PhoneNumber,
		MembershipId: req.MembershipId,
	}

	updatedCustomer, err := uc.customerRepository.Update(customerUpdate)
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
