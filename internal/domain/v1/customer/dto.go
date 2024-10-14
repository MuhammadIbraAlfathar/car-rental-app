package customerV1

import "time"

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Nik         string `json:"nik" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type CreateCustomerResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Nik         string    `json:"nik"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
