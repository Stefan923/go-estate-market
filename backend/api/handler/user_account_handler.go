package handler

import (
	"backend/api/dto"
	"backend/api/response"
	"backend/config"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAccountHandler struct {
	userAccountService *service.UserAccountService
}

func NewUserAccountHandler(config *config.Config) *UserAccountHandler {
	return &UserAccountHandler{
		userAccountService: service.NewUserAccountService(config),
	}
}

func (handler UserAccountHandler) Login(context *gin.Context) {
	request := new(dto.LoginRequest)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	tokenDetail, err := handler.userAccountService.Login(request)
	if err != nil {
		context.AbortWithStatusJSON(
			response.TranslateErrorToStatusCode(err),
			response.GenerateResponseWithError(nil, false, err))
		return
	}

	context.JSON(http.StatusCreated, response.GenerateResponse(tokenDetail, true))
}

func (handler UserAccountHandler) Register(context *gin.Context) {
	request := new(dto.RegisterRequest)
	err := context.ShouldBindJSON(request)
	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			response.GenerateResponseWithValidationError(nil, false, err))
		return
	}

	tokenDetail, err := handler.userAccountService.Register(context, request)
	if err != nil {
		context.AbortWithStatusJSON(
			response.TranslateErrorToStatusCode(err),
			response.GenerateResponseWithError(nil, false, err))
		return
	}

	context.JSON(http.StatusCreated, response.GenerateResponse(tokenDetail, false))
}
