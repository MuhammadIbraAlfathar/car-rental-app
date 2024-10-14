package bookingV1

import (
	"errors"
	carV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/car"
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	"time"
)

type UseCase struct {
	bookingRepository Repository
	carRepository     carV1.Repository
}

func NewUseCase(bookingRepository Repository, carRepository carV1.Repository) *UseCase {
	return &UseCase{
		bookingRepository: bookingRepository,
		carRepository:     carRepository,
	}
}

func (uc *UseCase) CreateBooking(req *CreateBookingRequest) (*v1Schema.Booking, error) {
	carId := req.CarId
	car, err := uc.carRepository.FindById(carId)
	if err != nil {
		return nil, errors.New("car not found")
	}
	carDailyRent := car.DailyRent

	startRent, err := time.Parse("2006-01-02", req.StartRent)
	if err != nil {
		return nil, errors.New("format start_date must YYYY-MM-DD")
	}

	endRent, err := time.Parse("2006-01-02", req.EndRent)
	if err != nil {
		return nil, errors.New("format end_rent must YYYY-MM-DD")
	}

	difference := endRent.Sub(startRent)
	differenceDays := int(difference.Hours() / 24)
	if differenceDays < 0 {
		differenceDays = -differenceDays
	}

	totalRent := differenceDays * carDailyRent

	bookingEntity := v1Schema.Booking{
		CustomerId: req.CustomerId,
		CarId:      req.CarId,
		StartRent:  startRent,
		EndRent:    endRent,
		TotalCost:  totalRent,
		Finished:   true,
	}

	newBooking, err := uc.bookingRepository.Create(&bookingEntity)
	if err != nil {
		return nil, err
	}

	return newBooking, nil
}

func (uc *UseCase) GetAllBooking() ([]*v1Schema.Booking, error) {
	bookings, err := uc.bookingRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
