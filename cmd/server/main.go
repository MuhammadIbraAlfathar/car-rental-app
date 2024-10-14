package main

import (
	"github.com/MuhammadIbraAlfathar/car-rental-app/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//test

	godotenv.Load()
	config.LoadEnv()
	_, err := config.NewPostgres()
	if err != nil {
		log.Println("ERROR TO CONNECT DATABASE")
	}
	r.Run()
}
