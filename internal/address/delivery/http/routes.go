package http

import "github.com/gin-gonic/gin"

func (h *handler) MapRoutes(r gin.RouterGroup) {
	r.GET("", h.GetCustomerAddresses())
	r.GET("/:address_id", h.GetAddress())
	r.POST("", h.CreateAddress())
	r.PATCH("/:address_id", h.UpdateAddress())
	r.DELETE("/:address_id", h.DeleteAddress())
}
