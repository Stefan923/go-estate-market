package router

import (
	"github.com/Stefan923/go-estate-market/api/handler"
	"github.com/Stefan923/go-estate-market/config"
	"github.com/gin-gonic/gin"
)

func StartAuthRouter(router *gin.RouterGroup, config *config.Config) {
	userAccountHandler := handler.NewUserAccountHandler(config)

	router.POST("/login", userAccountHandler.Login)
	router.POST("/register", userAccountHandler.Register)
}

func StartPropertiesRouter(router *gin.RouterGroup, config *config.Config) {
	propertyHandler := handler.NewPropertyHandler()

	router.GET("/:category/:pageNumber/:pageSize/:sortBy/:sortType", propertyHandler.GetAllByCategory)
	router.POST("/", propertyHandler.CreateProperty)
}
