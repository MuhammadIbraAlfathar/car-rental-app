package v2

import (
	"gorm.io/gorm"
	"time"
)

type Membership struct {
	Id             int            `json:"id" gorm:"primaryKey;autoIncrement"`
	MembershipName string         `json:"membership_name" gorm:"type:varchar(230);not null"`
	Discount       int            `json:"discount" gorm:"not null"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
