package handler

import (
	response2 "github.com/Stefan923/go-estate-market/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseHandler struct {
}

func (handler *BaseHandler) respondWithBadStatus(context *gin.Context, err error) {
	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		response2.GenerateResponseWithValidationError(nil, false, err))
}
