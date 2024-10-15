package config

import (
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	v2 "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewPostgres() (*gorm.DB, error) {
	host := Env.PostgresHost
	user := Env.PostgresUser
	password := Env.PostgresPassword
	dbName := Env.PostgresDbName
	port := Env.PostgresPort

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	log.Println("SUCCESS CONNECT TO DATABASE")

	//Migrate Table
	err = migratePostgresTable(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migratePostgresTable(db *gorm.DB) error {
	models := []interface{}{
		&v1Schema.Customer{},
		&v1Schema.Car{},
		&v1Schema.Booking{},

		//V2
		&v2.CustomerNew{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
