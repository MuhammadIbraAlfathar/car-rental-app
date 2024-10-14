package carV1

import v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"

type UseCase struct {
	carRepository Repository
}

func NewUseCase(carRepository Repository) *UseCase {
	return &UseCase{
		carRepository: carRepository,
	}
}

func (uc *UseCase) CreateCar(req *CreateCarRequest) (*v1Schema.Car, error) {
	carEntity := v1Schema.Car{
		Name:      req.Name,
		Stock:     req.Stock,
		DailyRent: req.DailyRent,
	}

	newCar, err := uc.carRepository.Create(&carEntity)
	if err != nil {
		return nil, err
	}

	return newCar, nil
}
