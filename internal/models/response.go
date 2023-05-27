package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fail response

func BadRequestResponse(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Code:   http.StatusBadRequest,
		Status: "bad request",
		Errors: errors,
		Data:   nil,
	})
}

func BadGatewayResponse(ctx *gin.Context, errors []string) {
	ctx.JSON(http.StatusBadGateway, Response{
		Code:   http.StatusBadGateway,
		Status: "bad gateway",
		Errors: errors,
		Data:   nil,
	})
}

// Success Response

func CreatedResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, Response{
		Code:   http.StatusCreated,
		Status: "created",
		Errors: []string{},
		Data:   data,
	})
}

func OkResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: "ok",
		Errors: []string{},
		Data:   data,
	})
}
