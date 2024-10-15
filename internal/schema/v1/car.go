package schema

import (
	"gorm.io/gorm"
	"time"
)

type Car struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"type:varchar(230);not null"`
	Stock     int            `json:"stock" gorm:"not null"`
	DailyRent int            `json:"daily_rent" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
