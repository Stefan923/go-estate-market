package router

import (
	"backend/api/handler"
	"backend/config"
	"github.com/gin-gonic/gin"
)

func StartAuthRouter(router *gin.RouterGroup, config *config.Config) {
	userAccountHandler := handler.NewUserAccountHandler(config)

	router.POST("/login", userAccountHandler.Login)
	router.POST("/register", userAccountHandler.Register)
}
