package carV1

import (
	"errors"
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
)

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

func (uc *UseCase) GetAllCar() ([]*v1Schema.Car, error) {
	car, err := uc.carRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (uc *UseCase) UpdateCar(carId int, req *UpdateCarRequest) (*v1Schema.Car, error) {
	car, err := uc.carRepository.FindById(carId)
	if err != nil {
		return nil, errors.New("car not found")
	}
	car.Name = req.Name
	car.Stock = req.Stock
	car.DailyRent = req.DailyRent

	updatedCustomer, err := uc.carRepository.Update(car)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (uc *UseCase) GetCarById(carId int) (*v1Schema.Car, error) {
	car, err := uc.carRepository.FindById(carId)
	if err != nil {
		return nil, errors.New("car not found")
	}

	return car, nil

}

func (uc *UseCase) DeleteCar(carId int) error {
	car, err := uc.carRepository.FindById(carId)
	if err != nil {
		return errors.New("car not found")
	}

	err = uc.carRepository.Delete(car)
	if err != nil {
		return err
	}

	return nil
}
