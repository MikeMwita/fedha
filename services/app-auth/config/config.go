package config

import (
	"fmt"
	"os"
)

//we will be implementing Redis||JWT||Db structs

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

type Jwt struct {
	Secret            string
	ExpiryMinutes     int
	RefreshExpiryDays int
}

type Redis struct {
	Port     string
	Host     string
	User     string
	Password string
}

type Config struct {
	Database     Database
	Redis        Redis
	MigrationUrl string
}

func LoadConfig() (*Config, error) {
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_HOST")
	}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_PORT")
	}

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_USER")
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_PASSWORD")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable DB_NAME")
	}

	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable REDIS_HOST")
	}

	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable OTP_HOST")
	}

	redisUser, ok := os.LookupEnv("REDIS_USER")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable REDIS_USER")
	}

	redisPassword, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("missing required environment variable REDIS_PASSWORD")
	}

	cfg := &Config{
		Database: Database{
			Port:     dbPort,
			Host:     dbHost,
			User:     dbUser,
			Dbname:   dbName,
			Password: dbPassword,
		},
		Redis: Redis{
			Port:     redisPort,
			Host:     redisHost,
			User:     redisUser,
			Password: redisPassword,
		},
	}
	return cfg, nil
}
