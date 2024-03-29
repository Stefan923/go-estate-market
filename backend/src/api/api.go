package api

import (
	"fmt"
	"github.com/Stefan923/go-estate-market/api/middleware"
	"github.com/Stefan923/go-estate-market/api/router"
	validator2 "github.com/Stefan923/go-estate-market/api/validator"
	"github.com/Stefan923/go-estate-market/config"
	"github.com/Stefan923/go-estate-market/metrics"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

func StartServer(config *config.Config) {
	gin.SetMode(config.Server.RunningMode)
	engine := gin.New()

	engine.Use(middleware.NewCorsMiddleware(config), middleware.NewPrometheusMiddleware())

	registerRoutes(engine, config)
	registerValidators()
	registerPrometheus()

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
		propertiesRoute := v1Route.Group("/properties")

		router.StartAuthRouter(userAccountsRoute, config)
		router.StartPropertiesRouter(propertiesRoute, config)
	}

	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
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

func registerPrometheus() {
	err := prometheus.Register(metrics.DatabaseCallCounter)
	if err != nil {
		if err != nil {
			log.Println("Error while registering prometheus metric: ", err)
		}
	}

	err = prometheus.Register(metrics.HttpDurationHistogram)
	if err != nil {
		if err != nil {
			log.Println("Error while registering prometheus metric: ", err)
		}
	}
}
