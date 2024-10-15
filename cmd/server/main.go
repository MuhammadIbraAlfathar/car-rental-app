package main

import (
	"github.com/MuhammadIbraAlfathar/car-rental-app/config"
	bookingV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/booking"
	carV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/car"
	customerV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/customer"
	bookingV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/booking"
	customerV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/customer"
	driverV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/driver"
	driverIncentiveV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/driver_incentive"
	membershipV2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v2/membership"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	//test

	godotenv.Load()
	config.LoadEnv()

	db, err := config.NewPostgres()
	if err != nil {
		log.Println("ERROR TO CONNECT DATABASE")
	}

	//CUSTOMER-V1
	customerRepo := customerV1.NewRepository(db)
	customerUseCase := customerV1.NewUseCase(customerRepo)
	customerV1.NewController(r, customerUseCase)

	//CAR
	carRepo := carV1.NewRepository(db)
	carUseCase := carV1.NewUseCase(carRepo)
	carV1.NewController(r, carUseCase)

	//BOOKING
	bookingRepo := bookingV1.NewRepository(db)
	bookingUseCase := bookingV1.NewUseCase(bookingRepo, carRepo)
	bookingV1.NewController(r, bookingUseCase)

	//CUSTOMER V2
	customerRepoV2 := customerV2.NewRepository(db)
	customerUseCaseV2 := customerV2.NewUseCase(customerRepoV2)
	customerV2.NewController(r, customerUseCaseV2)

	//MEMBERSHIP
	membershipRepository := membershipV2.NewRepository(db)

	//DRIVER
	driverRepository := driverV2.NewRepository(db)

	//DriverIncentive
	driverIncentiveRepo := driverIncentiveV2.NewRepository(db)

	//BOOKING-V2
	bookingRepoV2 := bookingV2.NewRepository(db)
	bookingUseCaseV2 := bookingV2.NewUseCase(bookingRepoV2, membershipRepository, driverRepository, driverIncentiveRepo, carRepo, customerRepoV2)
	bookingV2.NewController(r, bookingUseCaseV2)

	r.Run()
}
