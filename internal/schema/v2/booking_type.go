package v2

import (
	"gorm.io/gorm"
	"time"
)

type BookingType struct {
	Id              int            `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingTypeName string         `json:"booking_type_name" gorm:"type:varchar(230);not null"`
	Description     string         `json:"description" gorm:"type:text;not null"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
