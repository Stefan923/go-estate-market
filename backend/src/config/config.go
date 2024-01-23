package config

import (
	"errors"
	"github.com/spf13/viper"
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
}

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	configReader, err := LoadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error while loading config: %v", err)
	}

	config, err := ParseConfig(configReader)
	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.Server.ExternalPort = envPort
		log.Printf("Set external port from environment -> %s", config.Server.ExternalPort)
	} else {
		config.Server.ExternalPort = config.Server.InternalPort
		log.Printf("Set external port from environment -> %s", config.Server.ExternalPort)
	}
	if err != nil {
		log.Fatalf("Error while parsing config: %v", err)
	}

	return config
}

func ParseConfig(configReader *viper.Viper) (*Config, error) {
	var config Config
	err := configReader.Unmarshal(&config)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &config, nil
}

func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	configReader := viper.New()
	configReader.SetConfigType(fileType)
	configReader.SetConfigName(filename)
	configReader.AddConfigPath(".")
	configReader.AutomaticEnv()

	err := configReader.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, errors.New("could not find configuration file")
		}
		return nil, err
	}
	return configReader, nil
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
