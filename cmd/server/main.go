package main

import (
	"github.com/MuhammadIbraAlfathar/car-rental-app/config"
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

	r.Run()
}
