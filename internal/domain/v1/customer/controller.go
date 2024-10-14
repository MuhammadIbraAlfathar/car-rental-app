package customerV1

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

	customerGroup := engine.Group("/v1/customer")
	{
		customerGroup.POST("", controller.CreateCustomer())
	}
}

func (c *Controller) CreateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateCustomerRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			response.NewResponse(http.StatusBadRequest, err.Error(), "error").Send(ctx)
			return
		}

		newCustomer, err := c.uc.CreateCustomer(&req)
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}

		customerResponse := &CreateCustomerResponse{
			Id:          newCustomer.Id,
			Name:        newCustomer.Name,
			Nik:         newCustomer.Nik,
			PhoneNumber: newCustomer.PhoneNumber,
			CreatedAt:   newCustomer.CreatedAt,
			UpdatedAt:   newCustomer.UpdatedAt,
		}

		response.NewResponse(http.StatusCreated, "Success create customer", customerResponse).Send(ctx)
	}
}
