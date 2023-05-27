package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/dimy-tech-test/internal/address"
	"github.com/mrizalr/dimy-tech-test/internal/models"
)

type handler struct {
	usecase address.Usecase
}

func New(usecase address.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetCustomerAddresses() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Query("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		response, err := h.usecase.GetCustomerAddresses(customerID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) GetAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get address id
		addressID, err := strconv.Atoi(ctx.Param("address_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid address_id"})
			return
		}

		response, err := h.usecase.GetAddressByID(addressID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) CreateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Query("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		// Get payload
		payload := models.CustomerAddress{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		payload.CustomerID = customerID
		response, err := h.usecase.CreateAddress(payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.CreatedResponse(ctx, response)
	}
}

func (h *handler) UpdateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get address id
		addressID, err := strconv.Atoi(ctx.Param("address_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid address_id"})
			return
		}

		// Get payload
		payload := models.CustomerAddress{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.UpdateAddress(addressID, payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)

	}
}

func (h *handler) DeleteAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get address id
		addressID, err := strconv.Atoi(ctx.Param("address_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid address_id"})
			return
		}

		err = h.usecase.DeleteAddress(addressID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, gin.H{
			"message": fmt.Sprintf("address with id %v successfully deleted", addressID),
		})

	}
}
