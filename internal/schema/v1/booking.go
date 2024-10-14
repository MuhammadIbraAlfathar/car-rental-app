package schema

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	Id         int            `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerId int            `json:"customer_id" gorm:"not null"`
	Customer   Customer       `json:"customer"`
	CarId      int            `json:"car_id" gorm:"not null"`
	Car        Car            `json:"car"`
	StartRent  time.Time      `json:"start_rent" gorm:"type:date;not null"`
	EndRent    time.Time      `json:"end_rent" gorm:"type:date;not null"`
	TotalCost  int            `json:"total_cost" gorm:"not null"`
	Finished   bool           `json:"finished"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
