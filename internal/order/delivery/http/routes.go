package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(r *gin.RouterGroup) {
	r.GET("", h.GetOrders())
	r.GET("/:order_id", h.GetOrder())
	r.POST("/", h.CreateOrder())
	r.PATCH("/:order_id", h.UpdateOrder())
	r.DELETE("/:order_id", h.DeleteOrder())
}
