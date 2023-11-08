package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	return cfg, nil
}
