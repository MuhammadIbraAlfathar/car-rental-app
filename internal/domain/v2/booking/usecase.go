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
