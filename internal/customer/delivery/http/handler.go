package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/dimy-tech-test/internal/customer"
	"github.com/mrizalr/dimy-tech-test/internal/models"
)

type handler struct {
	usecase customer.Usecase
}

func New(usecase customer.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetCustomers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := h.usecase.GetCustomers()
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) GetCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Param("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		response, err := h.usecase.GetCustomerByID(customerID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) CreateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get payload
		payload := models.Customer{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.CreateCustomer(payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.CreatedResponse(ctx, response)
	}
}

func (h *handler) UpdateCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Param("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		// Get payload
		payload := models.Customer{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.UpdateCustomer(customerID, payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)

	}
}

func (h *handler) DeleteCustomer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Param("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		err = h.usecase.DeleteCustomer(customerID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, gin.H{
			"message": fmt.Sprintf("customer with id %v successfully deleted", customerID),
		})

	}
}
