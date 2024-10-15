package bookingV2

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

	bookingGroup := engine.Group("/v2/booking")
	{
		bookingGroup.POST("", controller.CreateBooking())
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
