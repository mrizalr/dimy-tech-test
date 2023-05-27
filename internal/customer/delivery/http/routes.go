package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(r *gin.RouterGroup) {
	r.GET("", h.GetCustomers())
	r.GET("/:customer_id", h.GetCustomer())
	r.POST("/", h.CreateCustomer())
	r.PATCH("/:customer_id", h.UpdateCustomer())
	r.DELETE("/:customer_id", h.DeleteCustomer())
}
