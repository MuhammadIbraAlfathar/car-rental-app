package bookingV2

import "time"

type CreateBookingRequest struct {
	CustomerId    int    `json:"customer_id" binding:"required"`
	CarId         int    `json:"car_id" binding:"required"`
	StartRent     string `json:"start_rent" binding:"required"`      // Format: YYYY-MM-DD
	EndRent       string `json:"end_rent" binding:"required"`        // Format: YYYY-MM-DD
	BookingTypeId int    `json:"booking_type_id" binding:"required"` //Value must 1 or 2
	DriverId      int    `json:"driver_id"`
	Finished      *bool  `json:"finished"`
}

type CreateBookingResponse struct {
	Id              int       `json:"id"`
	CustomerId      int       `json:"customer_id"`
	CarId           int       `json:"car_id"`
	StartRent       string    `json:"start_rent"` // Format: YYYY-MM-DD
	EndRent         string    `json:"end_rent"`   // Format: YYYY-MM-DD
	TotalCost       int       `json:"total_cost"`
	Discount        int       `json:"discount"`
	TotalDriverCost int       `json:"total_driver_cost"`
	Finished        *bool     `json:"finished"`
	CreatedAt       time.Time `json:"created_at"`
}
