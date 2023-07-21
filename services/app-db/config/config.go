package config

import (
	"github.com/MikeMwita/fedha.git/services/app-db/errors"
	"os"
)

type DbConfig struct {
	Host       string
	Port       string
	DbUsername string
	DbPassword string
}

func DbDnsConn() (string, error) {
	_, err := newDbConfig()
	if err != nil {
		return "", err
	}
	return "", nil
}

func newDbConfig() (*DbConfig, error) {
	getPort, ok := os.LookupEnv("POSTGRES_DB_PORT")
	if !ok {
		return nil, errors.ErrDBPortNotSet
	}
	getDbUsername, ok := os.LookupEnv("POSTGRES_DB_USERNAME")
	if !ok {
		return nil, errors.ErrDBUsernameNotSet
	}

	getDbPassword, ok := os.LookupEnv("POSTGRES_DB_PSWD")
	if !ok {
		return nil, errors.ErrDbPasswordNotSet
	}

	getDbHost, ok := os.LookupEnv("POSTGRES_DB_HOST")
	if !ok {
		return nil, errors.ErrDBHostNotSet
	}

	return &DbConfig{
		Host:       getDbHost,
		Port:       getPort,
		DbUsername: getDbUsername,
		DbPassword: getDbPassword,
	}, nil
}
