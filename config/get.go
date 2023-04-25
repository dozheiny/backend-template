package config

import (
	"github.com/joho/godotenv"
	"os"
	"runtime"
	"strings"
)

// Get will get key of environment and return value of key.
// if value is not exist, return error
func Get(key string) (string, error) {
	var value string

	value = os.Getenv(key)
	if strings.Compare(value, "") != 0 {
		return value, nil
	}

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	operationSystem := runtime.GOOS
	switch operationSystem {
	case "windows":
		if err := godotenv.Load(path + "\\.env"); err != nil {
			return "", err
		}

	case "linux":
		if err := godotenv.Load(path + "/.env"); err != nil {
			return "", err
		}
	}

	value = os.Getenv(key)

	if value == "" {
		return "", notFound
	}

	return value, nil
}
