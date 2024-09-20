package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvFromFile() error {
	environment, err := GetEnvVar("GO_ENV", "development")
	if err != nil {
		return err
	}

	if environment == "development" {
		err := godotenv.Load("./.env")
		if err != nil {
			return err
		}
	}

	return nil
}

func GetEnvVar(key string, defaultValue string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists && defaultValue == "" {
		return "", errors.New("environment variable not found")
	}

	if !exists && defaultValue != "" {
		return defaultValue, nil
	}

	return value, nil
}
