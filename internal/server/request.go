package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	idParam          = "id"
	originHeaderName = "Origin"
)

func GetParamFromRoute(c *gin.Context, name string) (int64, error) {
	routeId := c.Param(name)
	if routeId == "" {
		return 0, ErrInvalidIdParam
	}

	id, err := strconv.ParseInt(routeId, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w:%w", ErrConvertStringToNumber, err)
	}

	return id, nil
}

func GetIdFromRoute(c *gin.Context) (int64, error) {
	return GetParamFromRoute(c, idParam)
}

func GetInputFromBody[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.BindJSON(&inp); err != nil {
		return inp, ErrInvalidInputBody
	}

	return inp, nil
}

func GetInputFromQuery[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.BindQuery(&inp); err != nil {
		return inp, ErrInvalidInputQuery
	}

	return inp, nil
}

func GetInputFromForm[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.Bind(&inp); err != nil {
		return inp, ErrInvalidInputQuery
	}

	return inp, nil
}

func GetHost(c *gin.Context) string {
	return c.Request.Header.Get(originHeaderName)
}
