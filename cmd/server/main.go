package main

import (
	"github.com/MuhammadIbraAlfathar/car-rental-app/config"
	bookingV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/booking"
	carV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/car"
	customerV1 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/domain/v1/customer"
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

	//CUSTOMER
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

	r.Run()
}
