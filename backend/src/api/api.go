package api

import (
	"fmt"
	"github.com/Stefan923/go-estate-market/api/middleware"
	"github.com/Stefan923/go-estate-market/api/router"
	validator2 "github.com/Stefan923/go-estate-market/api/validator"
	"github.com/Stefan923/go-estate-market/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
)

func StartServer(config *config.Config) {
	gin.SetMode(config.Server.RunningMode)
	engine := gin.New()

	engine.Use(middleware.CreateCorsMiddleware(config))

	registerRoutes(engine, config)
	registerValidators()

	err := engine.Run(fmt.Sprintf(":%s", config.Server.InternalPort))
	if err != nil {
		fmt.Println("Error while starting server...")
	}
}

func registerRoutes(engine *gin.Engine, config *config.Config) {
	apiRoute := engine.Group("/api")

	v1Route := apiRoute.Group("/v1")
	{
		userAccountsRoute := v1Route.Group("/auth")

		router.StartAuthRouter(userAccountsRoute, config)
	}
}

func registerValidators() {
	validatorEngine, success := binding.Validator.Engine().(*validator.Validate)
	if success {
		err := validatorEngine.RegisterValidation("password", validator2.PasswordValidator, true)
		if err != nil {
			log.Println("Error while registering password validator: ", err)
		}
	}
}
