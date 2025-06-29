package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dataResponse[T any] struct {
	IsSuccess bool    `json:"isSuccess" binding:"required"`
	Data      *T      `json:"data,omitempty"`
	Message   *string `json:"message,omitempty"`
} // @name DataResponse

func SuccessResponse[T any](c *gin.Context, data *T) {
	c.JSON(http.StatusOK, createSuccessDataResponse(data))
}

func ErrorResponse[T any](c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		createErrorDataResponse[any](message))
}

func UnauthorizedResponse(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		createErrorDataResponse[any](message))
}

func createSuccessDataResponse[T any](data *T) dataResponse[T] {
	return dataResponse[T]{
		IsSuccess: true,
		Data:      data,
		Message:   nil,
	}
}

func createErrorDataResponse[T any](message string) dataResponse[T] {
	return dataResponse[T]{
		IsSuccess: false,
		Data:      nil,
		Message:   &message,
	}
}
