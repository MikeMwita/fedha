package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Database Database
	Jwt      Jwt
	Redis    Redis
	Metrics  Metrics
	Jaeger   Jaeger
	Server   ServerConfig
	Postgres PostgresConfig
	Session  Session
	Cookie   Cookie
	Logger   Logger

	JaegerCollectorHost string
	Env                 string
	IsTracingEnabled    bool
	ServiceName         string
	UseJaeger           bool
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

//	type DatabaseService struct {
//		Port string
//		Host string
//	}
type Jwt struct {
	Secret            string
	ExpiryMinutes     int
	RefreshExpiryDays int
}

type Redis struct {
	User     string
	Host     string
	Port     string
	Password string

	RedisAddr      string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	DB             int
}

type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

type Session struct {
	Prefix string
	Name   string
	Expire int
}
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	MaxConnectionIdle time.Duration
	Timeout           time.Duration
	MaxConnectionAge  time.Duration
	Time              time.Duration
}

type Metrics struct {
	URL         string
	ServiceName string
}

type Jaeger struct {
	JaegerCollectorHost string
	Host                string
	ServiceName         string
	LogSpans            bool
}

type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

type Logger struct {
	ctx               context.Context
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		slog.Error("unable to decode into struct")
		return nil, err
	}

	return &c, nil
}

func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadConf() (*Config, error) {
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
