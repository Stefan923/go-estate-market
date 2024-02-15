package handler

import (
	"github.com/Stefan923/go-estate-market/api/dto"
	response2 "github.com/Stefan923/go-estate-market/api/response"
	"github.com/Stefan923/go-estate-market/data/pagination"
	"github.com/Stefan923/go-estate-market/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PropertyHandler struct {
	propertyService service.PropertyService
}

func (handler *PropertyHandler) GetAllByCategory(context *gin.Context, category string, pageNumber int, pageSize int, sortBy string, sortType string) {
	authDetail, err := handler.propertyService.GetAllByCategory(category, &pagination.PageInfo{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		SortBy:     sortBy,
		SortType:   sortType,
	})
	if err != nil {
		context.AbortWithStatusJSON(
			response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}

	context.JSON(http.StatusOK, response2.GenerateResponse(authDetail, true))
}

func (handler *PropertyHandler) CreateProperty(context *gin.Context) {
	request := new(dto.PropertyCreationDto)
	err := context.ShouldBindJSON(request)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			response2.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	createdProperty, err := handler.propertyService.Save(context, request)
	if err != nil {
		context.AbortWithStatusJSON(
			response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}

	context.JSON(http.StatusCreated, response2.GenerateResponse(createdProperty, false))
}
