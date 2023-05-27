package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/order"
)

type handler struct {
	usecase order.Usecase
}

func New(usecase order.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := h.usecase.GetOrders()
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) GetOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Param("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		response, err := h.usecase.GetOrderByID(customerID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		customerID, err := strconv.Atoi(ctx.Query("customer_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid customer_id"})
			return
		}

		// Get Payload
		payload := models.OrderPayload{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}
		payload.CustomerID = customerID

		response, err := h.usecase.CreateOrder(payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.CreatedResponse(ctx, response)
	}
}

func (h *handler) UpdateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get order id
		orderID, err := strconv.Atoi(ctx.Param("order_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid order_id"})
			return
		}

		// Get payload
		payload := models.OrderPayload{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.UpdateOrder(orderID, payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)

	}
}

func (h *handler) DeleteOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get order id
		orderID, err := strconv.Atoi(ctx.Param("order_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid order_id"})
			return
		}

		err = h.usecase.DeleteOrder(orderID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, gin.H{
			"message": fmt.Sprintf("order with id %v successfully deleted", orderID),
		})

	}
}
