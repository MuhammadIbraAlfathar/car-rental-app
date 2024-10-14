package config

import (
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
	return db, nil
}
