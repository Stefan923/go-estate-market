package config

import "time"

type Config struct {
	Database DatabaseConfig
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
