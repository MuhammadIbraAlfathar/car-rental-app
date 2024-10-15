package bookingV2

import (
	"errors"
	carV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/car"
	customerV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/customer"
	driverV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/driver"
	driverIncentiveV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/driver_incentive"
	membershipV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/membership"
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"time"
)

type UseCase struct {
	bookingRepository         Repository
	carRepository             carV1.Repository
	customerRepository        customerV2.Repository
	membershipRepository      membershipV2.Repository
	driverRepository          driverV2.Repository
	driverIncentiveRepository driverIncentiveV2.Repository
}

func NewUseCase(bookingRepository Repository, membershipRepository membershipV2.Repository, driverRepository driverV2.Repository, driverIncentiveRepository driverIncentiveV2.Repository, carRepository carV1.Repository, customerRepository customerV2.Repository) *UseCase {
	return &UseCase{
		bookingRepository:         bookingRepository,
		membershipRepository:      membershipRepository,
		driverRepository:          driverRepository,
		driverIncentiveRepository: driverIncentiveRepository,
		customerRepository:        customerRepository,
		carRepository:             carRepository,
	}
}

func (uc *UseCase) CreateBooking(req *CreateBookingRequest) (*v2.BookingNew, error) {

	//Calculate total rent
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

	//Calculate Discount
	customerId := req.CustomerId
	customer, err := uc.customerRepository.FindById(customerId)
	customerMembershipId := customer.MembershipId
	discount := 0

	if customerMembershipId == 1 {
		discount = (totalRent * 4) / 100
	} else if customerMembershipId == 2 {
		discount = (totalRent * 7) / 100
	} else if customerMembershipId == 3 {
		discount = (totalRent * 15) / 100
	} else if customerMembershipId == 0 {
		discount = 0
	} else {
		return nil, errors.New("error id membership")
	}

	//TOTAL DRIVER COST CALCULATE
	totalDriverCost := 0
	driverIncentive := 0
	driverId := req.DriverId

	//validate booking type and driver
	if driverId == 0 && req.BookingTypeId == 1 {
		totalDriverCost = 0
		driverIncentive = 0
	} else if driverId != 0 && req.BookingTypeId == 2 {
		driver, err := uc.driverRepository.FindById(driverId)
		if err != nil {
			return nil, errors.New("driver not found")
		}
		driverDailyCost := driver.DailyCost
		driverIncentive = (totalRent * 5) / 100
		totalDriverCost = driverDailyCost * differenceDays
	}

	bookingEntity := v2.BookingNew{
		CustomerId:      customerId,
		CarId:           carId,
		StartRent:       startRent,
		EndRent:         endRent,
		TotalCost:       totalRent,
		Finished:        true,
		Discount:        discount,
		BookingTypeId:   req.BookingTypeId,
		DriverId:        driverId,
		TotalDriverCost: totalDriverCost,
	}

	booking, err := uc.bookingRepository.Create(&bookingEntity)
	bookingId := booking.Id

	driverIncentiveEntity := &v2.DriverIncentive{
		BookingId: bookingId,
		Incentive: driverIncentive,
	}

	_, err = uc.driverIncentiveRepository.Create(driverIncentiveEntity)
	if err != nil {
		return nil, err
	}

	return booking, nil

}

func (uc *UseCase) GetAllBooking() ([]*v2.BookingNew, error) {
	bookings, err := uc.bookingRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (uc *UseCase) GetBookingByCustomerId(customerId int) ([]*v2.BookingNew, error) {
	bookings, err := uc.bookingRepository.FindByCustomerId(customerId)
	if err != nil {
		return nil, errors.New("data booking not found")
	}

	return bookings, nil

}

func (uc *UseCase) UpdateBooking(bookID int, req *CreateBookingRequest) (*v2.BookingNew, error) {
	booking, err := uc.bookingRepository.FindById(bookID)
	if err != nil {
		return nil, errors.New("book not found")
	}

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

	//Calculate Discount
	customerId := req.CustomerId
	customer, err := uc.customerRepository.FindById(customerId)
	customerMembershipId := customer.MembershipId
	discount := 0

	if customerMembershipId == 1 {
		discount = (totalRent * 4) / 100
	} else if customerMembershipId == 2 {
		discount = (totalRent * 7) / 100
	} else if customerMembershipId == 3 {
		discount = (totalRent * 15) / 100
	} else if customerMembershipId == 0 {
		discount = 0
	} else {
		return nil, errors.New("error id membership")
	}

	//TOTAL DRIVER COST CALCULATE
	totalDriverCost := 0
	driverIncentiveCost := 0
	driverId := req.DriverId

	//validate booking type and driver
	if driverId == 0 && req.BookingTypeId == 1 {
		totalDriverCost = 0
		driverIncentiveCost = 0
	} else if driverId != 0 && req.BookingTypeId == 2 {
		driver, err := uc.driverRepository.FindById(driverId)
		if err != nil {
			return nil, errors.New("driver not found")
		}
		driverDailyCost := driver.DailyCost
		driverIncentiveCost = (totalRent * 5) / 100
		totalDriverCost = driverDailyCost * differenceDays
	}

	updateBooking := v2.BookingNew{
		Id:              booking.Id,
		CustomerId:      req.CustomerId,
		CarId:           req.CarId,
		StartRent:       startRent,
		EndRent:         endRent,
		TotalCost:       totalRent,
		Finished:        true,
		Discount:        discount,
		BookingTypeId:   req.BookingTypeId,
		DriverId:        req.DriverId,
		TotalDriverCost: totalDriverCost,
	}

	updatedBooking, err := uc.bookingRepository.Update(&updateBooking)
	if err != nil {
		return nil, err
	}

	driverIncentive, err := uc.driverIncentiveRepository.FindByBookingId(updatedBooking.Id)
	if err != nil {
		return nil, err
	}

	driverIncentiveUpdate := &v2.DriverIncentive{
		Id:        driverIncentive.Id,
		BookingId: updatedBooking.Id,
		Incentive: driverIncentiveCost,
	}

	_, err = uc.driverIncentiveRepository.Update(driverIncentiveUpdate)
	if err != nil {
		return nil, err
	}

	return updatedBooking, nil
}

func (uc *UseCase) DeleteBooking(bookId int) error {
	booking, err := uc.bookingRepository.FindById(bookId)
	if err != nil {
		return errors.New("booking not found")
	}

	err = uc.bookingRepository.Delete(booking)
	if err != nil {
		return err
	}

	return nil
}
