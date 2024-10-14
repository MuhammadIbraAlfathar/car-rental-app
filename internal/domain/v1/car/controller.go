package carV1

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

	carGroup := engine.Group("/v1/car")
	{
		carGroup.POST("", controller.CreateCar())
	}
}

func (c *Controller) CreateCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateCarRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		newCar, err := c.uc.CreateCar(&req)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}

		carResponse := &CreateCustomerResponse{
			Id:        newCar.Id,
			Name:      newCar.Name,
			Stock:     newCar.Stock,
			DailyRent: newCar.DailyRent,
			CreatedAt: newCar.CreatedAt,
			UpdatedAt: newCar.UpdatedAt,
		}

		response.NewResponse(http.StatusCreated, "Success create car", carResponse).Send(ctx)
	}
}
