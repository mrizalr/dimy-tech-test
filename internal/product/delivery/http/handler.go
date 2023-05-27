package http

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/dimy-tech-test/internal/models"
	"github.com/mrizalr/dimy-tech-test/internal/product"
)

type handler struct {
	usecase product.Usecase
}

func New(usecase product.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := h.usecase.GetProducts()
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) GetProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get product id
		productID, err := strconv.Atoi(ctx.Param("product_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid product_id"})
			return
		}

		response, err := h.usecase.GetProductByID(productID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)
	}
}

func (h *handler) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get payload
		payload := models.Product{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.CreateProduct(payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.CreatedResponse(ctx, response)
	}
}

func (h *handler) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get product id
		productID, err := strconv.Atoi(ctx.Param("product_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid product_id"})
			return
		}

		// Get payload
		payload := models.Product{}
		if err := ctx.ShouldBindJSON(&payload); err != nil {
			models.BadRequestResponse(ctx, []string{err.Error()})
			return
		}

		response, err := h.usecase.UpdateProduct(productID, payload)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, response)

	}
}

func (h *handler) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get customer id
		productID, err := strconv.Atoi(ctx.Param("product_id"))
		if err != nil {
			models.BadRequestResponse(ctx, []string{"invalid product_id"})
			return
		}

		err = h.usecase.DeleteProduct(productID)
		if err != nil {
			models.BadGatewayResponse(ctx, []string{err.Error()})
			return
		}
		models.OkResponse(ctx, gin.H{
			"message": fmt.Sprintf("product with id %v successfully deleted", productID),
		})

	}
}
