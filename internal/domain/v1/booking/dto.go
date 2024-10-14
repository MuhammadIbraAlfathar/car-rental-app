package bookingV1

import "time"

type CreateBookingRequest struct {
	CustomerId int    `json:"customer_id" binding:"required"`
	CarId      int    `json:"car_id" binding:"required"`
	StartRent  string `json:"start_rent" binding:"required"` // Format: YYYY-MM-DD
	EndRent    string `json:"end_rent" binding:"required"`   // Format: YYYY-MM-DD
	Finished   *bool  `json:"finished"`
}

type BookingResponse struct {
	Id         int                     `json:"id"`
	CustomerId int                     `json:"customer_id"`
	Customer   CustomerBookingResponse `json:"customer"`
	CarId      int                     `json:"car_id"`
	Car        CarResponse             `json:"car"`
	StartRent  string                  `json:"start_rent"` // Format: YYYY-MM-DD
	EndRent    string                  `json:"end_rent"`   // Format: YYYY-MM-DD
	TotalCost  int                     `json:"total_cost"`
	Finished   *bool                   `json:"finished"`
	CreatedAt  time.Time               `json:"created_at"`
}

type CustomerBookingResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type CarResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Stock     int    `json:"stock"`
	DailyRent int    `json:"daily_rent"`
}
