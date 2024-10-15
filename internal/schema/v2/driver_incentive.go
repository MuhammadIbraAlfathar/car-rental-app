package v2

import (
	"gorm.io/gorm"
	"time"
)

type DriverIncentive struct {
	Id        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	BookingId int            `json:"booking_id" gorm:"not null"`
	Incentive int            `json:"incentive" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
