package customerV1

import v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"

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
