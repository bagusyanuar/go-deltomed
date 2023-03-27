package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configurationImpl struct{}

func (config *configurationImpl) Get(key string) string {
	return os.Getenv(key)
}

func New() Config {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}
	return &configurationImpl{}
}
