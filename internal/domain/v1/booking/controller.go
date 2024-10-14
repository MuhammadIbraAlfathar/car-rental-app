package bookingV1

import (
	"github.com/MuhammadIbraAlfathar/car-rental-app/internal/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	uc *UseCase
}

func NewController(engine *gin.Engine, uc *UseCase) {
	controller := &Controller{uc: uc}

	bookingGroup := engine.Group("/v1/booking")
	{
		bookingGroup.POST("", controller.CreateCar())
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

		//carResponse := &CreateCarResponse{
		//	Id:        newCar.Id,
		//	Name:      newCar.Name,
		//	Stock:     newCar.Stock,
		//	DailyRent: newCar.DailyRent,
		//	CreatedAt: newCar.CreatedAt,
		//	UpdatedAt: newCar.UpdatedAt,
		//}

		bookResponse := &BookingResponse{
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
