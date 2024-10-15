package bookingV2

import (
	"github.com/MuhammadIbraAlfathar/car-rental-app/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	uc *UseCase
}

func NewController(engine *gin.Engine, uc *UseCase) {
	controller := &Controller{uc: uc}

	bookingGroup := engine.Group("/v2/booking")
	{
		bookingGroup.POST("", controller.CreateBooking())
		bookingGroup.GET("", controller.GetAllBooking())
		bookingGroup.GET("/:id", controller.GetBookingByCustomerId())
	}

}

func (c *Controller) CreateBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateBookingRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		newBooking, err := c.uc.CreateBooking(&req)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}
		bookResponse := &CreateBookingResponse{
			Id:              newBooking.Id,
			CustomerId:      newBooking.CustomerId,
			CarId:           newBooking.CarId,
			StartRent:       newBooking.StartRent.Format("2006-01-02"),
			EndRent:         newBooking.EndRent.Format("2006-01-02"),
			TotalCost:       newBooking.TotalCost,
			Discount:        newBooking.Discount,
			TotalDriverCost: newBooking.TotalDriverCost,
			Finished:        &newBooking.Finished,
			CreatedAt:       newBooking.CreatedAt,
		}

		response.NewResponse(http.StatusCreated, "Success create booking", bookResponse).Send(ctx)
	}
}

func (c *Controller) GetAllBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bookings, err := c.uc.GetAllBooking()
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}

		var bookingResponses []*BookingResponse
		for _, i := range bookings {
			bookResponse := &BookingResponse{
				Id:         i.Id,
				CustomerId: i.CustomerId,
				Customer: CustomerBookingResponse{
					Id:           i.Customer.Id,
					Name:         i.Customer.Name,
					PhoneNumber:  i.Customer.PhoneNumber,
					MembershipId: i.Customer.MembershipId,
				},
				CarId: i.CarId,
				Car: CarResponse{
					Id:        i.Car.Id,
					Name:      i.Car.Name,
					Stock:     i.Car.Stock,
					DailyRent: i.Car.DailyRent,
				},
				StartRent:       i.StartRent.Format("2006-01-02"),
				EndRent:         i.EndRent.Format("2006-01-02"),
				TotalCost:       i.TotalCost,
				Discount:        i.Discount,
				BookTypeId:      i.BookingTypeId,
				DriverId:        i.DriverId,
				TotalDriverCost: i.TotalDriverCost,
				Finished:        &i.Finished,
				CreatedAt:       i.CreatedAt,
			}

			bookingResponses = append(bookingResponses, bookResponse)
		}

		response.NewResponse(http.StatusOK, "Success get data booking", bookingResponses).Send(ctx)
	}
}

func (c *Controller) GetBookingByCustomerId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		customerId, _ := strconv.Atoi(id)
		bookings, err := c.uc.GetBookingByCustomerId(customerId)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}

		var bookingResponses []*BookingResponseByUserId
		for _, i := range bookings {
			bookResponse := &BookingResponseByUserId{
				Id:         i.Id,
				CustomerId: i.CustomerId,
				CarId:      i.CarId,
				Car: CarResponse{
					Id:        i.Car.Id,
					Name:      i.Car.Name,
					Stock:     i.Car.Stock,
					DailyRent: i.Car.DailyRent,
				},
				StartRent:       i.StartRent.Format("2006-01-02"),
				EndRent:         i.EndRent.Format("2006-01-02"),
				TotalCost:       i.TotalCost,
				Discount:        i.Discount,
				BookTypeId:      i.BookingTypeId,
				DriverId:        i.DriverId,
				TotalDriverCost: i.TotalDriverCost,
				Finished:        &i.Finished,
				CreatedAt:       i.CreatedAt,
			}

			bookingResponses = append(bookingResponses, bookResponse)
		}

		response.NewResponse(http.StatusOK, "Success get data booking", bookingResponses).Send(ctx)
	}
}
