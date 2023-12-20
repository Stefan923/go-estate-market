package config

import "time"

type Config struct {
	Database DatabaseConfig
	JWT      JwtConfig
	Auth     AuthConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	DbName             string
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

type AuthConfig struct {
	BCryptCost int
}

type ServerConfig struct {
	RunningMode  string
	InternalPort string
	ExternalPort string
}
