package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

// App Config
type Config struct {
	Server ServerConfig
}

// Fiber Config
type ServerConfig struct {
	Host string `env:"FIBER_HOST,required"`
	Port string `env:"FIBER_PORT,required"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		panic(err)
	}

	fmt.Println("[Config] environments has been loaded.")

	return cfg
}
