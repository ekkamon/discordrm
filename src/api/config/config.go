package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	_ "github.com/joho/godotenv/autoload"
)

// App Config
type Config struct {
	Server ServerConfig
	PgSQL  PostgreSQLConfig
	JWT    JWTConfig
}

// Fiber Config
type ServerConfig struct {
	Host string `env:"FIBER_HOST,required"`
	Port string `env:"FIBER_PORT,required"`
}

// PostgteSQL Config
type PostgreSQLConfig struct {
	Host     string `env:"PG_HOST,required"`
	Port     int    `env:"PG_PORT,required"`
	Username string `env:"PG_USER,required"`
	Password string `env:"PG_PASSWORD,required"`
	DBName   string `env:"PG_DBNAME,required"`
	SSLMode  string `env:"PG_SSLMODE,required"`
	Timezone string `env:"PG_TIMEZONE,required"`
}

// JWT Config
type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET_KEY,required"`
}

func LoadConfig() *Config {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		panic(err)
	}

	fmt.Println("[Config] environments has been loaded.")

	return cfg
}
