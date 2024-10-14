package carV1

import "time"

type CreateCarRequest struct {
	Name      string `json:"name" binding:"required"`
	Stock     int    `json:"stock" binding:"required"`
	DailyRent int    `json:"daily_rent" binding:"required"`
}

type CreateCustomerResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	DailyRent int       `json:"daily_rent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
