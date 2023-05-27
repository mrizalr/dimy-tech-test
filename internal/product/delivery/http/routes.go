package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(r *gin.RouterGroup) {
	r.GET("", h.GetProducts())
	r.GET("/:product_id", h.GetProduct())
	r.POST("", h.CreateProduct())
	r.PATCH("/:product_id", h.UpdateProduct())
	r.DELETE("/:product_id", h.DeleteProduct())
}
