package customerV2

import "time"

type CreateCustomerRequest struct {
	Name         string `json:"name" binding:"required"`
	Nik          string `json:"nik" binding:"required"`
	MembershipId int    `json:"membership_id"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
}

type CreateCustomerResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Nik          string    `json:"nik"`
	PhoneNumber  string    `json:"phone_number"`
	MembershipId int       `json:"membership_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdatedCustomerRequest struct {
	Name         string `json:"name"`
	Nik          string `json:"nik"`
	MembershipId int    `json:"membership_id"`
	PhoneNumber  string `json:"phone_number"`
}
