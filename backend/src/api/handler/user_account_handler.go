package handler

import (
	"github.com/Stefan923/go-estate-market/api/dto"
	response2 "github.com/Stefan923/go-estate-market/api/response"
	"github.com/Stefan923/go-estate-market/config"
	"github.com/Stefan923/go-estate-market/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAccountHandler struct {
	BaseHandler
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
		handler.respondWithBadStatus(context, err)
		return
	}

	authDetail, err := handler.userAccountService.Login(request)
	if err != nil {
		context.AbortWithStatusJSON(
			response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}

	context.JSON(http.StatusCreated, response2.GenerateResponse(authDetail, true))
}

func (handler UserAccountHandler) Register(context *gin.Context) {
	request := new(dto.RegisterRequest)
	err := context.ShouldBindJSON(request)
	if err != nil {
		handler.respondWithBadStatus(context, err)
		return
	}

	authDetail, err := handler.userAccountService.Register(context, request)
	if err != nil {
		context.AbortWithStatusJSON(
			response2.TranslateErrorToStatusCode(err),
			response2.GenerateResponseWithError(nil, false, err))
		return
	}

	context.JSON(http.StatusCreated, response2.GenerateResponse(authDetail, false))
}
