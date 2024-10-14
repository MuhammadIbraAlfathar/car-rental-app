package carV1

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

	carGroup := engine.Group("/v1/car")
	{
		carGroup.POST("", controller.CreateCar())
		carGroup.GET("", controller.GetAllCar())
		carGroup.PUT("/:id", controller.UpdateCar())
		carGroup.GET("/:id", controller.GetCarById())
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

		carResponse := &CreateCarResponse{
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

func (c *Controller) GetAllCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		car, err := c.uc.GetAllCar()
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}

		var carResponses []*CreateCarResponse
		for _, i := range car {
			carResponse := &CreateCarResponse{
				Id:        i.Id,
				Name:      i.Name,
				Stock:     i.Stock,
				DailyRent: i.DailyRent,
				CreatedAt: i.CreatedAt,
				UpdatedAt: i.UpdatedAt,
			}

			carResponses = append(carResponses, carResponse)
		}

		response.NewResponse(http.StatusOK, "Success get data car", carResponses).Send(ctx)
	}
}

func (c *Controller) UpdateCar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		carId, err := strconv.Atoi(id)
		if err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		var req UpdateCarRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		updatedCar, err := c.uc.UpdateCar(carId, &req)
		if err != nil {
			response.NewResponse(http.StatusNotFound, err.Error(), "error").Send(ctx)
			return
		}

		updatedCarResponse := &CreateCarResponse{
			Id:        updatedCar.Id,
			Name:      updatedCar.Name,
			Stock:     updatedCar.Stock,
			DailyRent: updatedCar.DailyRent,
			CreatedAt: updatedCar.CreatedAt,
			UpdatedAt: updatedCar.UpdatedAt,
		}

		response.NewResponse(http.StatusOK, "Success update car", updatedCarResponse).Send(ctx)
	}
}

func (c *Controller) GetCarById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		carId, err := strconv.Atoi(id)
		if err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		car, err := c.uc.GetCarById(carId)
		if err != nil {
			response.NewResponse(http.StatusNotFound, err.Error(), "something went wrong").Send(ctx)
			return
		}

		carResponse := &CreateCarResponse{
			Id:        car.Id,
			Name:      car.Name,
			Stock:     car.Stock,
			DailyRent: car.DailyRent,
			CreatedAt: car.CreatedAt,
			UpdatedAt: car.UpdatedAt,
		}

		response.NewResponse(http.StatusOK, "Success get car by id", carResponse).Send(ctx)
	}
}
