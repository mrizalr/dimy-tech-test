package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/payment"
)

type handler struct {
	usecase payment.Usecase
}

func New(usecase payment.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetPaymentMethods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := h.usecase.GetPaymentMethods()
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) GetPaymentMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get product id
		paymentID, err := strconv.Atoi(ctx.Param("payment_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid payment_id"})
			return
		}

		response, err := h.usecase.GetPaymentMethodByID(paymentID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) CreatePaymentMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get payload
		payload := models.PaymentMethod{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.CreatePaymentMethod(payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.CreatedResponse(ctx, response)
	}
}

func (h *handler) UpdatePaymentMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get product id
		paymentID, err := strconv.Atoi(ctx.Param("payment_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid payment_id"})
			return
		}

		// Get payload
		payload := models.PaymentMethod{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.UpdatePaymentMethod(paymentID, payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)

	}
}

func (h *handler) DeletePaymentMethod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		paymentID, err := strconv.Atoi(ctx.Param("payment_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid payment_id"})
			return
		}

		err = h.usecase.DeletePaymentMethod(paymentID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, gin.H{
			"message": fmt.Sprintf("payment method with id %v successfully deleted", paymentID),
		})

	}
}
