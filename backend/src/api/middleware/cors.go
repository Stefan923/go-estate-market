package middleware

import (
	"github.com/Stefan923/go-estate-market/config"
	"github.com/gin-gonic/gin"
)

func CreateCorsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Server.Cors.AllowedOrigins)
		context.Header("Access-Control-Allow-Credentials", cfg.Server.Cors.AllowCredentials)
		context.Header("Access-Control-Allow-Headers", cfg.Server.Cors.AllowedHeaders)
		context.Header("Access-Control-Allow-Methods", cfg.Server.Cors.AllowedMethods)
		context.Header("Access-Control-Max-Age", cfg.Server.Cors.MaxAge)
		context.Set("content-type", cfg.Server.Cors.ContentType)
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}
