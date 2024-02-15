package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Database DatabaseConfig
	Auth     AuthConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	DatabaseName       string
	SSLMode            string
	MaxIdleConnections int
	MaxOpenConnections int
	ConnMaxLifetime    time.Duration
}

type JwtConfig struct {
	AccessTokenExpireDurationMinutes  time.Duration
	RefreshTokenExpireDurationMinutes time.Duration
	AccessTokenSecret                 string
	RefreshTokenSecret                string
}

type PasswordConfig struct {
	BCryptCost       int
	IncludeChars     bool
	IncludeDigits    bool
	IncludeUppercase bool
	IncludeLowercase bool
	MinLength        int
	MaxLength        int
}

type AuthConfig struct {
	Password PasswordConfig
	JWT      JwtConfig
}

type ServerConfig struct {
	RunningMode  string
	InternalPort string
	ExternalPort string
	Cors         CorsConfig
}

type CorsConfig struct {
	AllowedOrigins   string
	AllowedHeaders   string
	AllowedMethods   string
	AllowCredentials string
	ContentType      string
	MaxAge           string
}

func GetConfig() *Config {
	fileReader := FileReader[Config]{}
	configFilePath := getConfigPath(os.Getenv("APP_ENV"))

	config := fileReader.GetContent(configFilePath)
	if config == nil {
		return nil
	}

	environmentPort := os.Getenv("PORT")
	if environmentPort != "" {
		config.Server.ExternalPort = environmentPort
		log.Printf("Set external port from environment -> %s", config.Server.ExternalPort)
	} else {
		config.Server.ExternalPort = config.Server.InternalPort
		log.Printf("Set external port from environment -> %s", config.Server.ExternalPort)
	}

	return config
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/app/config/docker-config"
	} else if env == "production" {
		return "/app/config/production-config"
	} else {
		return "/app/config/development-config"
	}
}
