package v2

import (
	v1Schema "github.com/MuhammadIbraAlfathar/car-rental-app/internal/schema/v1"
	"gorm.io/gorm"
	"time"
)

type BookingNew struct {
	Id              int            `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerId      int            `json:"customer_id" gorm:"not null"`
	Customer        CustomerNew    `json:"customer"`
	CarId           int            `json:"car_id" gorm:"not null"`
	Car             v1Schema.Car   `json:"car"`
	StartRent       time.Time      `json:"start_rent" gorm:"type:date;not null"`
	EndRent         time.Time      `json:"end_rent" gorm:"type:date;not null"`
	TotalCost       int            `json:"total_cost" gorm:"not null"`
	Finished        bool           `json:"finished"`
	Discount        int            `json:"discount"`
	BookingTypeId   int            `json:"booking_type_id" gorm:"not null"`
	DriverId        int            `json:"driver_id"`
	TotalDriverCost int            `json:"total_driver_cost"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
