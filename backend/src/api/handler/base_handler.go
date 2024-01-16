package handler

import (
	"context"
	response2 "github.com/Stefan923/go-estate-market/api/response"
	error3 "github.com/Stefan923/go-estate-market/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create[T any, R any](context *gin.Context, caller func(context context.Context, request *T) (*R, error)) {
	request := new(T)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			response2.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	responseObject, err := caller(context, request)
	if err != nil {
		context.AbortWithStatusJSON(response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusCreated, response2.GenerateResponse(responseObject, true))
}

func Update[T any, R any](context *gin.Context, caller func(context context.Context, id int, request *T) (*R, error)) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	request := new(T)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			response2.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	responseObject, err := caller(context, id, request)
	if err != nil {
		context.AbortWithStatusJSON(response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusOK, response2.GenerateResponse(responseObject, true))
}

func Delete(context *gin.Context, caller func(context context.Context, id int) error) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound,
			response2.GenerateResponseWithError(nil, false, &error3.InternalError{
				EndUserMessage: error3.InvalidPathParameter,
			}))
		return
	}

	err := caller(context, id)
	if err != nil {
		context.AbortWithStatusJSON(response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusOK, response2.GenerateResponse(nil, true))
}

func GetById[R any](context *gin.Context, caller func(context context.Context, id int) (*R, error)) {
	id, _ := strconv.Atoi(context.Params.ByName("id"))
	if id == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound,
			response2.GenerateResponse(nil, false))
		return
	}

	responseObject, err := caller(context, id)
	if err != nil {
		context.AbortWithStatusJSON(response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}
	context.JSON(http.StatusOK, response2.GenerateResponse(responseObject, true))
}
