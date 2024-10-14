package schema

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Id          int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(230);not null"`
	Nik         string         `json:"nik" gorm:"type:varchar(230);not null"`
	PhoneNumber string         `json:"phone_number" gorm:"type:varchar(230);not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
