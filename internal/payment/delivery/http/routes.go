package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(r *gin.RouterGroup) {
	r.GET("", h.GetPaymentMethods())
	r.GET("/:payment_id", h.GetPaymentMethod())
	r.POST("", h.CreatePaymentMethod())
	r.PATCH("/:payment_id", h.UpdatePaymentMethod())
	r.DELETE("/:payment_id", h.DeletePaymentMethod())
}
