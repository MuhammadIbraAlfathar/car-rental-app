package bookingV1

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

	bookingGroup := engine.Group("/v1/booking")
	{
		bookingGroup.POST("", controller.CreateCar())
		bookingGroup.GET("", controller.GetAllBooking())
		bookingGroup.PUT("/:id", controller.UpdateBooking())
		bookingGroup.GET("/:id", controller.GetBookingById())
		bookingGroup.DELETE("/:id", controller.DeleteBooking())
	}
}

func (c *Controller) CreateCar() gin.HandlerFunc {
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
			Id:         newBooking.Id,
			CustomerId: newBooking.CustomerId,
			CarId:      newBooking.CarId,
			StartRent:  newBooking.StartRent.Format("2006-01-02"),
			EndRent:    newBooking.EndRent.Format("2006-01-02"),
			TotalCost:  newBooking.TotalCost,
			Finished:   &newBooking.Finished,
			CreatedAt:  newBooking.CreatedAt,
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
					Id:          i.Customer.Id,
					Name:        i.Customer.Name,
					PhoneNumber: i.Customer.PhoneNumber,
				},
				CarId: i.CarId,
				Car: CarResponse{
					Id:        i.Car.Id,
					Name:      i.Car.Name,
					Stock:     i.Car.Stock,
					DailyRent: i.Car.DailyRent,
				},
				StartRent: i.StartRent.Format("2006-01-02"),
				EndRent:   i.EndRent.Format("2006-01-02"),
				TotalCost: i.TotalCost,
				Finished:  &i.Finished,
				CreatedAt: i.CreatedAt,
			}

			bookingResponses = append(bookingResponses, bookResponse)
		}

		response.NewResponse(http.StatusOK, "Success get data booking", bookingResponses).Send(ctx)
	}
}

func (c *Controller) UpdateBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		bookId, err := strconv.Atoi(id)
		if err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		var req UpdatedBookingRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		updatedBook, err := c.uc.UpdateBooking(bookId, &req)
		if err != nil {
			response.NewResponse(http.StatusNotFound, err.Error(), "error").Send(ctx)
			return
		}

		bookResponse := &UpdateBookingResponse{
			Id:         updatedBook.Id,
			CustomerId: updatedBook.CustomerId,
			CarId:      updatedBook.CarId,
			StartRent:  updatedBook.StartRent.Format("2006-01-02"),
			EndRent:    updatedBook.EndRent.Format("2006-01-02"),
			TotalCost:  updatedBook.TotalCost,
			Finished:   &updatedBook.Finished,
			CreatedAt:  updatedBook.CreatedAt,
			UpdatedAt:  updatedBook.UpdatedAt,
		}

		response.NewResponse(http.StatusOK, "Success update booking", bookResponse).Send(ctx)
	}
}

func (c *Controller) GetBookingById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		bookId, err := strconv.Atoi(id)
		if err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		booking, err := c.uc.GetBookingById(bookId)
		if err != nil {
			response.NewResponse(http.StatusNotFound, err.Error(), "something went wrong").Send(ctx)
			return
		}

		bookResponse := &GetBookingByIdResponse{
			Id:         booking.Id,
			CustomerId: booking.CustomerId,
			CarId:      booking.CarId,
			Car: CarResponse{
				Id:        booking.Car.Id,
				Name:      booking.Car.Name,
				Stock:     booking.Car.Stock,
				DailyRent: booking.Car.DailyRent,
			},
			StartRent: booking.StartRent.Format("2006-01-02"),
			EndRent:   booking.EndRent.Format("2006-01-02"),
			TotalCost: booking.TotalCost,
			Finished:  &booking.Finished,
			CreatedAt: booking.CreatedAt,
		}
		response.NewResponse(http.StatusOK, "Success get booking by id", bookResponse).Send(ctx)
	}
}

func (c *Controller) DeleteBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		bookId, _ := strconv.Atoi(id)

		err := c.uc.DeleteBooking(bookId)
		if err != nil {
			response.NewResponse(http.StatusNotFound, err.Error(), "error").Send(ctx)
			return
		}

		response.NewResponse(http.StatusOK, "Success delete booking", "").Send(ctx)
	}
}
