package handler

import (
	"github.com/Stefan923/go-estate-market/api/dto"
	response2 "github.com/Stefan923/go-estate-market/api/response"
	"github.com/Stefan923/go-estate-market/data/pagination"
	"github.com/Stefan923/go-estate-market/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PropertyHandler struct {
	BaseHandler
	propertyService *service.PropertyService
}

func NewPropertyHandler() *PropertyHandler {
	return &PropertyHandler{
		propertyService: service.NewPropertyService(),
	}
}

func (handler *PropertyHandler) GetAllByCategory(context *gin.Context) {
	category := context.Param("category")
	pageNumber, err := strconv.Atoi(context.Param("pageNumber"))
	if err != nil {
		handler.respondWithBadStatus(context, err)
		return
	}
	pageSize, err := strconv.Atoi(context.Param("pageSize"))
	if err != nil {
		handler.respondWithBadStatus(context, err)
		return
	}
	sortBy := context.Param("sortBy")
	sortType := context.Param("sortType")

	authDetail, err := handler.propertyService.GetAllByCategory(category, &pagination.PageInfo{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		SortBy:     sortBy,
		SortType:   sortType,
	})
	if err != nil {
		handler.respondWithBadStatus(context, err)
		return
	}

	context.JSON(http.StatusOK, response2.GenerateResponse(authDetail, true))
}

func (handler *PropertyHandler) CreateProperty(context *gin.Context) {
	request := new(dto.PropertyCreationDto)
	err := context.ShouldBindJSON(request)
	if err != nil {
		handler.respondWithBadStatus(context, err)
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
