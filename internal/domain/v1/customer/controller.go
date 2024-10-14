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
		customerGroup.GET("", controller.GetAllCustomer())
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

func (c *Controller) GetAllCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer, err := c.uc.GetAllCustomer()
		if err != nil {
			response.NewResponse(http.StatusInternalServerError, err.Error(), "something went wrong").Send(ctx)
			return
		}

		var customerResponses []*CreateCustomerResponse
		for _, i := range customer {
			customerResponse := &CreateCustomerResponse{
				Id:          i.Id,
				Name:        i.Name,
				Nik:         i.Nik,
				PhoneNumber: i.PhoneNumber,
				CreatedAt:   i.CreatedAt,
				UpdatedAt:   i.UpdatedAt,
			}

			customerResponses = append(customerResponses, customerResponse)
		}

		response.NewResponse(http.StatusOK, "Success get data customer", customerResponses).Send(ctx)
	}
}
