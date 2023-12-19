package handler

import (
	"backend/api/response"
	error2 "backend/error"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create[T any, R any](context *gin.Context, caller func(context context.Context, request *T) (*R, error)) {
	request := new(T)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			response.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	responseObject, err := caller(context, request)
	if err != nil {
		context.AbortWithStatusJSON(response.TranslateErrorToStatusCode(err),
			response.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusCreated, response.GenerateResponse(responseObject, true))
}

func Update[T any, R any](context *gin.Context, caller func(context context.Context, id int, request *T) (*R, error)) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	request := new(T)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			response.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	responseObject, err := caller(context, id, request)
	if err != nil {
		context.AbortWithStatusJSON(response.TranslateErrorToStatusCode(err),
			response.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusOK, response.GenerateResponse(responseObject, true))
}

func Delete(context *gin.Context, caller func(context context.Context, id int) error) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound,
			response.GenerateResponseWithError(nil, false, &error2.InternalError{
				EndUserMessage: error2.InvalidPathParameter,
			}))
		return
	}

	err := caller(context, id)
	if err != nil {
		context.AbortWithStatusJSON(response.TranslateErrorToStatusCode(err),
			response.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusOK, response.GenerateResponse(nil, true))
}

func GetById[R any](context *gin.Context, caller func(context context.Context, id int) (*R, error)) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound,
			response.GenerateResponse(nil, false))
		return
	}

	responseObject, err := caller(context, id)
	if err != nil {
		context.AbortWithStatusJSON(response.TranslateErrorToStatusCode(err),
			response.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusOK, response.GenerateResponse(responseObject, true))
}
